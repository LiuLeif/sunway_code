#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-06-14 18:11
from gensim.models import word2vec
import torch
from torch.utils.data import Dataset, DataLoader
import numpy as np

FEATURE_SIZE = 50
CHUNK_SIZE = 50
HIDDEN_SIZE = 50
EPOCH = 10000

print("training word2vec")
sentences = word2vec.Text8Corpus("xwz.txt")
vec = word2vec.Word2Vec(sentences, size=FEATURE_SIZE, iter=100, min_count=0)
vocabulary = vec.wv.vocab.keys()
char_to_idx = {}
idx_to_char = {}
for v in vocabulary:
    index = vec.wv.vocab[v].index
    char_to_idx[v] = index
    idx_to_char[index] = v

print("word2vec done")


class Converter:
    def __init__(self):
        self.char_to_idx = char_to_idx
        self.idx_to_char = idx_to_char

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
    conv = Converter()
    s = " ".join([conv.to_char(i) for i in indexs])
    print(s)


class WordSamplerDataset(Dataset):
    def __init__(self, ):
        self.conv = Converter()
        with open("xwz.txt", "r") as f:
            self.data = f.read()
            self.data = self.data.split()

    def __getitem__(self, index):
        start = np.random.randint(0, len(self.data) - CHUNK_SIZE)
        chunk_data = self.data[start:start + CHUNK_SIZE]
        inp = self.conv.char_tensor(chunk_data[:-1])
        target = self.conv.char_tensor(chunk_data[1:])
        return inp, target

    def __len__(self):
        return EPOCH


class WordSampler(torch.nn.Module):
    def __init__(self, hidden_size):
        super().__init__()
        self.encoder = torch.nn.Embedding(len(idx_to_char), FEATURE_SIZE)
        self.encoder.weight.data.copy_(torch.from_numpy(vec.wv.vectors))
        self.encoder.requires_grad = False

        self.rnn = torch.nn.LSTMCell(FEATURE_SIZE, hidden_size)
        self.fc = torch.nn.Linear(hidden_size, len(idx_to_char))

    def forward(self, inp, h, c):
        out = self.encoder(inp)
        h, c = self.rnn(out, (h, c))
        out = self.fc(h)
        return out, h, c


train_loader = DataLoader(WordSamplerDataset())
model = WordSampler(HIDDEN_SIZE)

optimizer = torch.optim.Adam(model.parameters())

criterion = torch.nn.CrossEntropyLoss()

epoch = 0
for X, Y in train_loader:
    model.zero_grad()
    h = torch.zeros(len(X), HIDDEN_SIZE)
    c = torch.zeros(len(X), HIDDEN_SIZE)
    loss = 0
    for i in range(X.size(1)):
        out, h, c = model(X[:, i], h, c)
        loss += criterion(out, Y[:, i].view(-1))

    loss.backward()
    optimizer.step()
    if epoch % 100 == 0:
        print("training #%d: loss %f" % (epoch, loss.item() / CHUNK_SIZE))
    epoch += 1
