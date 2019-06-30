#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-07-25 11:06
import numpy as np
import matplotlib.pyplot as plt
from sklearn.datasets import make_moons
from matplotlib.colors import ListedColormap
from torch.utils.data import Dataset, DataLoader
import torch

model = torch.nn.Sequential(
    torch.nn.Linear(2, 50), torch.nn.ReLU(), torch.nn.Linear(50, 1), torch.nn.Sigmoid()
)


class MoonDataset(Dataset):
    def __init__(self):
        X, Y = make_moons(n_samples=1000, noise=0.2)
        self.X = torch.from_numpy(X).float()
        self.Y = torch.from_numpy(Y).float().view(-1, 1)

    def __getitem__(self, index):
        return self.X[index], self.Y[index]

    def __len__(self):
        return len(self.X)


dataset = MoonDataset()
loader = DataLoader(dataset, batch_size=100)

criterion = torch.nn.BCELoss()

optimizer = torch.optim.Adam(model.parameters(), weight_decay=0.001)


def train():
    for i in range(200):
        for x, y in loader:
            loss = criterion(model(x), y)
            optimizer.zero_grad()
            loss.backward()
            optimizer.step()


def visualize():
    cm = ListedColormap(["#FF0000", "#0000FF"])
    plt.scatter(x=dataset.X[:, 0], y=dataset.X[:, 1], c=dataset.Y[:, 0], cmap=cm)
    xx, yy = np.meshgrid(np.arange(-4, 4, 0.02), np.arange(-4, 4, 0.02))
    X = np.c_[xx.ravel(), yy.ravel()]

    y_hat = model(torch.from_numpy(X).float()).detach().numpy()
    y_hat = (y_hat > 0.5)[:, 0]
    y_hat = y_hat.reshape(xx.shape)

    cm = ListedColormap(["#FF0000", "#0000FF"])
    plt.contour(xx, yy, y_hat, cmap=cm)
    plt.show()


train()
visualize()
