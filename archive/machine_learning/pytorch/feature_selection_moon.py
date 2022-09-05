#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-07-25 11:06
import numpy as np
import matplotlib.pyplot as plt
from sklearn.datasets import make_moons
from matplotlib.colors import ListedColormap
from torch.utils.data import Dataset, DataLoader
import torch

NOISES = 10
N_SAMPLES = 2000

model = torch.nn.Sequential(
    torch.nn.Linear(2 + NOISES, 10),
    torch.nn.BatchNorm1d(10),
    torch.nn.Tanh(),
    torch.nn.Linear(10, 1),
    torch.nn.Sigmoid(),
)


class MoonDataset(Dataset):
    def __init__(self):
        X, Y = make_moons(n_samples=N_SAMPLES, noise=0.2)
        r = np.random.randn(N_SAMPLES, NOISES)
        X = np.c_[(X, r)]
        self.X = torch.from_numpy(X).float()
        self.Y = torch.from_numpy(Y).float().view(-1, 1)

    def __getitem__(self, index):
        return self.X[index], self.Y[index]

    def __len__(self):
        return len(self.X)


training_set = MoonDataset()
training_loader = DataLoader(training_set, batch_size=100)

test_set = MoonDataset()
test_loader = DataLoader(test_set, batch_size=N_SAMPLES)

criterion = torch.nn.BCELoss()

optimizer = torch.optim.Adam(model.parameters(), weight_decay=0.001)


def train():
    model.train()
    for i in range(1000):
        for x, y in training_loader:
            loss = criterion(model(x), y)
            optimizer.zero_grad()
            loss.backward()
            optimizer.step()
        if i % 20 == 0:
            print(loss.item())


def test():
    model.eval()
    for x, y in training_loader:
        y_hat = model(x)
        y_hat = y_hat > 0.5
        accu = torch.sum(y_hat == y.byte()).item() / 100
        print("train:", accu)
        break

    for x, y in test_loader:
        y_hat = model(x)
        y_hat = y_hat > 0.5
        accu = torch.sum(y_hat == y.byte()).item() / N_SAMPLES
        print("test:", accu)


def visualize():
    plt.close()
    cm = ListedColormap(["#FF0000", "#0000FF"])
    plt.scatter(x=test_set.X[:, 0], y=test_set.X[:, 1], c=test_set.Y[:, 0], cmap=cm)

    xx, yy = np.meshgrid(np.arange(-4, 4, 0.02), np.arange(-4, 4, 0.02))
    X = np.c_[xx.ravel(), yy.ravel()]

    if NOISES != 0:
        y_hat = (
            model(
                torch.cat(
                    (torch.from_numpy(X).float(), torch.ones(X.shape[0], NOISES)), dim=1
                )
            )
            .detach()
            .numpy()
        )
    else:
        y_hat = model(torch.from_numpy(X).float()).detach().numpy()

    y_hat = (y_hat > 0.5)[:, 0]
    y_hat = y_hat.reshape(xx.shape)

    cm = ListedColormap(["#FF0000", "#0000FF"])
    plt.contour(xx, yy, y_hat, cmap=cm)
    plt.show()


train()
# visualize()
