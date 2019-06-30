#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-06-08 17:35
import torch
import torchvision

train_data = torchvision.datasets.MNIST(
    root="./mnist_data",
    train=True,
    transform=torchvision.transforms.ToTensor(),
    download=True,
)

test_data = torchvision.datasets.MNIST(
    root="./mnist_data",
    train=False,
    transform=torchvision.transforms.ToTensor(),
    download=True,
)

train_loader = torch.utils.data.DataLoader(
    dataset=train_data, batch_size=200, shuffle=True
)
test_loader = torch.utils.data.DataLoader(
    dataset=test_data, batch_size=len(test_data), shuffle=False
)


class CNN(torch.nn.Module):
    def __init__(self):
        super().__init__()
        self.conv1 = torch.nn.Sequential(
            torch.nn.Conv2d(1, 20, 5), torch.nn.ReLU(), torch.nn.MaxPool2d(2, 2)
        )
        self.fc = torch.nn.Linear(12 * 12 * 20, 10)

    def forward(self, input):
        out = self.conv1(input)
        out = self.fc(out.view(out.size(0), -1))
        return out


model = CNN()
criterion = torch.nn.CrossEntropyLoss()
optimizer = torch.optim.Adam(model.parameters())

for i in range(5):
    total_loss = 0
    for (x, y) in train_loader:
        out = model(x)
        loss = criterion(out, y)

        total_loss += loss.item()
        optimizer.zero_grad()
        loss.backward()
        optimizer.step()

    with torch.no_grad():
        for x, y in test_loader:
            predicted = model(x)
            correct = torch.max(predicted, 1)[1] == y

    print(
        "training: #%d, loss %f, accu %f"
        % (i, loss / len(train_loader), correct.sum() * 100 / len(test_data))
    )
