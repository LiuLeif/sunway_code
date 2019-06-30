#!/usr/bin/env python3
from matplotlib import pyplot as plt
from sklearn import datasets
from sklearn.decomposition import PCA
from sklearn.manifold import TSNE

iris = datasets.load_iris()

X = iris.data
y = iris.target

target_ids = range(len(iris.target_names))

# PCA
pca = PCA(n_components=2, random_state=0)
X_2d = pca.fit_transform(X)

plt.subplot(131)
plt.title("pca")
colors = 'r', 'g', 'b', 'c', 'm', 'y', 'k', 'w', 'orange', 'purple'
for i, c, label in zip(target_ids, colors, iris.target_names):
    plt.scatter(X_2d[y == i, 0], X_2d[y == i, 1], c=c, label=label)

# TSNE
tsne = TSNE(n_components=2, random_state=0)
X_2d = tsne.fit_transform(X)
colors = 'r', 'g', 'b', 'c', 'm', 'y', 'k', 'w', 'orange', 'purple'
plt.subplot(132)
plt.title("tsne")
for i, c, label in zip(target_ids, colors, iris.target_names):
    plt.scatter(X_2d[y == i, 0], X_2d[y == i, 1], c=c, label=label)

# autoencoder
import torch
from torch import nn
from torch import optim
from torch.utils.data import DataLoader, Dataset


# ---------- data ----------
class PlainDataset(Dataset):
    def __init__(self):
        self.X = torch.from_numpy(X).float()
        self.Y = self.X

    def __getitem__(self, index):
        return self.X[index], self.Y[index]

    def __len__(self):
        return len(self.X)


training_set = PlainDataset()
training_loader = DataLoader(training_set, batch_size=30, shuffle=True)


def train():
    for i in range(5000):
        for x, y in training_loader:
            loss = criterion(model(x), y)
            optimizer.zero_grad()
            loss.backward()
            optimizer.step()
        # if i % 20 == 0:
        #     print("epoch #%d: loss: %f" % (i, loss.item()))


# ---------- model ----------
model = nn.Sequential(nn.Linear(4, 2), nn.ReLU(), nn.Linear(2, 4))
criterion = nn.MSELoss()
optimizer = optim.Adam(model.parameters(), lr=1e-3, weight_decay=0.001)

train()
X_2d = model[0](torch.from_numpy(X).float()).detach().numpy()
colors = 'r', 'g', 'b', 'c', 'm', 'y', 'k', 'w', 'orange', 'purple'
plt.subplot(133)
plt.title("auto-encoder")
for i, c, label in zip(target_ids, colors, iris.target_names):
    plt.scatter(X_2d[y == i, 0], X_2d[y == i, 1], c=c, label=label)

plt.show()
