#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-06-11 16:11
import torch
from torch.utils.data import Dataset, DataLoader
import string
import numpy as np

CHUNK_SIZE = 200
HIDDEN_SIZE = 50
VOCABULARY = string.printable
EPOCH = 10000
LEARNNG_RATE = 0.01


class Converter:
    def __init__(self):
        self.char_to_idx = {}
        self.idx_to_char = {}
        for i, j in enumerate(VOCABULARY):
            self.char_to_idx[j] = i
            self.idx_to_char[i] = j

    def to_idx(self, c):
        return self.char_to_idx[c]

    def to_char(self, i):
        return self.idx_to_char[i]

    def char_tensor(self, string):
        tensor = torch.zeros(len(string)).long()
        for c in range(len(string)):
            tensor[c] = self.to_idx(string[c])
        return tensor


def sample():
    h = torch.zeros(1, HIDDEN_SIZE)
    c = torch.zeros(1, HIDDEN_SIZE)
    inp = torch.zeros(1).long()

    indexs = []
    with torch.no_grad():
        for i in range(1000):
            out, h, c = model(inp, h, c)
            out = torch.nn.functional.softmax(out)
            selected = torch.multinomial(out, 1)
            indexs.append(selected.item())
            inp = selected.view(-1)
    conv = Converter()
    s = "".join([conv.to_char(i) for i in indexs])
    print(s)


class WordSamplerDataset(Dataset):
    def __init__(self,):
        self.conv = Converter()
        with open("data.txt", "r") as f:
            self.data = f.read()

    def __getitem__(self, index):
        start = np.random.randint(0, len(self.data) - CHUNK_SIZE)
        chunk_data = self.data[start : start + CHUNK_SIZE]
        inp = self.conv.char_tensor(chunk_data[:-1])
        target = self.conv.char_tensor(chunk_data[1:])
        return inp, target

    def __len__(self):
        return EPOCH


def save():
    torch.save(model.state_dict(), "word_sampler.pt")


def restore():
    model.load_state_dict(torch.load("word_sampler.pt"))
    print("load model from word_sampler.pt")


class WordSampler(torch.nn.Module):
    def __init__(self, hidden_size):
        super().__init__()
        self.encoder = torch.nn.Embedding(len(VOCABULARY), hidden_size)
        # self.rnn = torch.nn.RNNCell(hidden_size, hidden_size)
        self.rnn = torch.nn.LSTMCell(hidden_size, hidden_size)
        self.fc = torch.nn.Linear(hidden_size, len(VOCABULARY))

    def forward(self, inp, h, c):
        out = self.encoder(inp)
        h, c = self.rnn(out, (h, c))
        out = self.fc(h)
        return out, h, c


train_loader = DataLoader(WordSamplerDataset())
model = WordSampler(HIDDEN_SIZE)

try:
    restore()
except:
    pass
optimizer = torch.optim.Adam(model.parameters(), lr=LEARNNG_RATE)

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

# for i in range(10):
#     sample()
