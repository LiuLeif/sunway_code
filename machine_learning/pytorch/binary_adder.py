#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-06-11 16:11
import torch
from torch.utils.data import Dataset, DataLoader
import numpy as np

HIDDEN_SIZE = 20
BATCH_SIZE = 30
MAX_ORDER = 6


def predict(a, b):
    a = np.array([[a]], dtype=np.uint8)
    b = np.array([[b]], dtype=np.uint8)

    z = np.stack((a, b), axis=1)
    Z = np.unpackbits(z, axis=2).transpose(0, 2, 1)
    Z = np.flip(Z, 1)

    with torch.no_grad():
        hidden = torch.zeros(1, 1, HIDDEN_SIZE)
        out, _ = model(torch.from_numpy(Z.copy()).float(), hidden)
        ret = out.numpy().reshape(-1)
    ret = ret > 0.5
    ret = np.flip(ret, 0)
    ret = np.packbits(ret)[0]
    print(ret)


class BinaryAddDataset(Dataset):
    def __init__(self):
        largest_num = pow(2, MAX_ORDER)
        a = np.random.randint(0, largest_num / 2, dtype=np.uint8, size=(300, 1))
        b = np.random.randint(0, largest_num / 2, dtype=np.uint8, size=(300, 1))
        c = a + b

        C = np.unpackbits(c, axis=1)
        C = np.expand_dims(C, axis=2)

        z = np.stack((a, b), axis=1)
        Z = np.unpackbits(z, axis=2).transpose(0, 2, 1)

        Z = np.flip(Z, 1)
        C = np.flip(C, 1)

        self.Z = torch.from_numpy(Z.copy()).float()
        self.C = torch.from_numpy(C.copy()).float()

    def __getitem__(self, index):
        return self.Z[index], self.C[index]

    def __len__(self):
        return len(self.Z)


class BinaryAdder(torch.nn.Module):
    def __init__(self, hidden_size):
        super().__init__()
        self.rnn = torch.nn.RNN(2, hidden_size, batch_first=True)
        self.out = torch.nn.Linear(hidden_size, 1)

    def forward(self, x, hidden):
        hidden, _ = self.rnn(x, hidden)
        out = self.out(hidden)
        return out, hidden


def save():
    torch.save(model.state_dict(), "binary_adder.pt")


def restore():
    model.load_state_dict(torch.load("binary_adder.pt"))
    print("load model from binary_adder.pt")


train_loader = DataLoader(BinaryAddDataset(), batch_size=BATCH_SIZE)

model = BinaryAdder(HIDDEN_SIZE)
try:
    restore()
except:
    pass

optimizer = torch.optim.Adam(model.parameters())

criterion = torch.nn.MSELoss()

for epoch in range(5000):
    total_loss = 0
    for x, y in train_loader:
        hidden = torch.zeros(1, len(x), HIDDEN_SIZE)
        model.zero_grad()
        out, hidden = model(x, hidden)
        loss = criterion(out, y)
        loss.backward()
        total_loss += loss.item()
        optimizer.step()
    if epoch % 1000 == 0:
        print("training #%d, loss: %f" % (epoch, total_loss / len(train_loader)))

predict(123, 12)
