#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-07-31 17:42
import numpy as np
import matplotlib.pyplot as plt
from torch.utils.data import Dataset, DataLoader
import torch


def f(x):
    """ function to approximate by polynomial interpolation"""
    return x * np.sin(x)


class TestDataset(Dataset):
    def __init__(self):
        X = np.linspace(0, 10, 500)
        rng = np.random.RandomState(0)
        rng.shuffle(X)
        X = np.sort(X[:100])
        Y = f(X)
        self.X = torch.from_numpy(X).float().view(-1, 1)
        self.Y = torch.from_numpy(Y).float().view(-1, 1)

    def __getitem__(self, index):
        return self.X[index], self.Y[index]

    def __len__(self):
        return len(self.X)


dataset = TestDataset()


def visualize():
    X = dataset.X[:, np.newaxis]
    Y = dataset.Y[:, np.newaxis]
    x_plot = np.linspace(0, 10, 500)
    x_plot = x_plot[:, np.newaxis]

    lw = 2
    plt.plot(
        x_plot, f(x_plot), color="cornflowerblue", linewidth=lw, label="ground truth"
    )

    x_predict = torch.from_numpy(x_plot).float()
    y_predict = model(x_predict)

    plt.scatter(
        x_predict.detach().numpy(),
        y_predict.detach().numpy(),
        color="red",
        s=30,
        marker="x",
        label="predicted points",
    )

    plt.scatter(X, Y, color="navy", s=30, marker="o", label="training points")
    plt.show()


loader = DataLoader(dataset, batch_size=50)

model = torch.nn.Sequential(
    torch.nn.Linear(1, 100),
    torch.nn.ReLU(),
    torch.nn.Linear(100, 50),
    torch.nn.ReLU(),
    torch.nn.Linear(50, 1),
)

criterion = torch.nn.MSELoss()
optimizer = torch.optim.Adam(model.parameters(), lr=1e-2)


def train():
    for i in range(10000):
        for x, y in loader:
            loss = criterion(model(x), y)
            optimizer.zero_grad()
            loss.backward()
            optimizer.step()
        if i % 100 == 0:
            print("epoch #%d: %f" % (i, loss.item()))


train()

visualize()
