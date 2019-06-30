#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-08-14 13:39
import numpy as np
import matplotlib.pyplot as plt

import torch
from torch import nn
from torch import optim
from torch.utils.data import DataLoader, Dataset
import torch.nn.functional as F


# ---------- data ----------
class PlainDataset(Dataset):
    def __init__(self):
        x = torch.randn(1000, 10)
        # x = torch.round(torch.rand(1000) * 200)
        # self.z = torch.round(torch.rand(1000) * 200)
        # x = x.unsqueeze(1)
        # x = torch.cat((x, x * 2, x * 3, x * 4, x * 5, x * 6, x * 7, x * 8,
        #                x * 9, x * 10), 1)
        # self.X = x + self.z.unsqueeze(1)
        self.X = x
        self.Y = self.X

    def __getitem__(self, index):
        return self.X[index], self.Y[index]

    def __len__(self):
        return len(self.X)


training_set = PlainDataset()
training_loader = DataLoader(training_set, batch_size=100, shuffle=True)


# ---------- helper ----------
def test():
    m = model[0]

    # x = torch.tensor([[2, 4, 6, 8, 10, 12, 14, 16, 18, 20]]).float()
    # x = x + 1
    # y_hat = model(x)
    # print("orig: ", x, " new: ", y_hat, "a:", m(x))

    # x = torch.tensor([[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]]).float()
    # x = x + 2
    # y_hat = model(x)
    # print("orig: ", x, " new: ", y_hat, "a:", m(x))

    # x = torch.tensor([[10, 20, 30, 40, 50, 60, 70, 80, 90, 100]]).float()
    # x = x + 3
    # y_hat = model(x)
    # print("orig: ", x, " new: ", y_hat, "a:", m(x))
    # torch.manual_seed(1000)
    x = torch.randn(1, 10)
    y_hat = model(x)
    print("orig: ", x, " new: ", y_hat, "a:", m(x))


def train():
    for i in range(300):
        for x, y in training_loader:
            loss = criterion(model(x), y)
            optimizer.zero_grad()
            loss.backward()
            optimizer.step()
        if i % 20 == 0:
            print("epoch #%d: loss: %f" % (i, loss.item()))


# ---------- model ----------

model = nn.Sequential(nn.Linear(10, 10), nn.ReLU(), nn.Linear(10, 10))
criterion = nn.MSELoss()
optimizer = optim.Adam(model.parameters())

train()
