#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-06-14 18:11
from gensim.models import word2vec
import torch
from torch.utils.data import Dataset, DataLoader
import numpy as np
import argparse
import os

FEATURE_SIZE = 50
CHUNK_SIZE = 20
HIDDEN_SIZE = 100
EPOCH = 200000


class Converter:
    def __init__(self):
        self._load_word2vec()
        print(f"token size: {self.size()}")

    def _load_word2vec(self):
        if os.path.exists("rr.w2v"):
            self.vec = word2vec.Word2Vec.load("rr.w2v")
        else:
            sentences = word2vec.Text8Corpus("rr.txt")
            self.vec = word2vec.Word2Vec(
                sentences, vector_size=FEATURE_SIZE, epochs=200, min_count=0
            )
            self.vec.save("rr.w2v")
        vocabulary = self.vec.wv.key_to_index
        self.char_to_idx = {}
        self.idx_to_char = {}
        for v, index in vocabulary.items():
            self.char_to_idx[v] = index
            self.idx_to_char[index] = v

    def get_vector(self):
        return torch.from_numpy(self.vec.wv.vectors)

    def size(self):
        return len(self.idx_to_char)

    def to_idx(self, c):
        return self.char_to_idx[c]

    def to_char(self, i):
        return self.idx_to_char[i]

    def char_tensor(self, s):
        tensor = torch.zeros(len(s)).long()
        for c in range(len(s)):
            tensor[c] = self.to_idx(s[c])
        return tensor


def sample():
    model = WordSampler()
    model.load_state_dict(torch.load("random_chinese.pt"))

    h = torch.zeros(1, HIDDEN_SIZE)
    c = torch.zeros(1, HIDDEN_SIZE)
    inp = torch.zeros(1).long()

    indexs = []
    with torch.no_grad():
        for i in range(200):
            out, h, c = model(inp, h, c)
            out = torch.nn.functional.softmax(out)
            selected = torch.multinomial(out, 1)
            indexs.append(selected.item())
            inp = selected.view(-1)
    s = " ".join([global_converter.to_char(i) for i in indexs])
    print(s)


class WordSamplerDataset(Dataset):
    def __init__(
        self,
    ):
        with open("rr.txt", "r") as f:
            self.data = f.read()
            self.data = self.data.split()

    def __getitem__(self, index):
        start = np.random.randint(0, len(self.data) - CHUNK_SIZE)
        chunk_data = self.data[start : start + CHUNK_SIZE]
        inp = global_converter.char_tensor(chunk_data[:-1])
        target = global_converter.char_tensor(chunk_data[1:])
        return inp, target

    def __len__(self):
        return EPOCH


class WordSampler(torch.nn.Module):
    def __init__(self):
        super().__init__()
        self.encoder = torch.nn.Embedding(global_converter.size(), FEATURE_SIZE)
        self.encoder.weight.data.copy_(global_converter.get_vector())
        self.encoder.requires_grad = False

        self.rnn = torch.nn.LSTMCell(FEATURE_SIZE, HIDDEN_SIZE)
        self.fc = torch.nn.Linear(HIDDEN_SIZE, global_converter.size())

    def forward(self, inp, h, c):
        out = self.encoder(inp)
        h, c = self.rnn(out, (h, c))
        out = self.fc(h)
        return out, h, c


def train():
    train_loader = DataLoader(WordSamplerDataset())
    model = WordSampler().to("cuda")
    optimizer = torch.optim.Adam(model.parameters())
    criterion = torch.nn.CrossEntropyLoss()

    best_loss = None
    epoch = 0
    for X, Y in train_loader:
        X = X.to("cuda")
        Y = Y.to("cuda")
        model.zero_grad()
        h = torch.zeros(len(X), HIDDEN_SIZE).to("cuda")
        c = torch.zeros(len(X), HIDDEN_SIZE).to("cuda")
        loss = 0
        for i in range(X.size(1)):
            out, h, c = model(X[:, i], h, c)
            loss += criterion(out, Y[:, i].view(-1))

        loss.backward()
        optimizer.step()
        if epoch % 100 == 0:
            if best_loss is None or best_loss > loss.item():
                best_loss = loss.item()
                print("training #%d: loss %f" % (epoch, loss.item() / CHUNK_SIZE))
                torch.save(model.state_dict(), "random_chinese.pt")

        epoch += 1


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--mode", choices=["train", "sample"], required=True)
    args = parser.parse_args()

    global global_converter
    global_converter = Converter()

    if args.mode == "train":
        train()
    elif args.mode == "sample":
        sample()
