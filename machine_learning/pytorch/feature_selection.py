#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-08-16 18:25
import torch
from sklearn import preprocessing

# torch.manual_seed(1000)
w1 = torch.randn(1)
w1.requires_grad = True
w2 = torch.randn(1)
w2.requires_grad = True

criterion = torch.nn.MSELoss()
optimizer = torch.optim.SGD([w1, w2], lr=1e-3)

# N_SAMPLES = 2
# x1 = torch.rand(N_SAMPLES, 1) + 1
# x2 = torch.rand(N_SAMPLES, 1) + 2


def calc(x1, x2):
    a = x1.prod() + x2.prod()
    b = (x1**2 + x2**2).prod()
    print(a**2 / b)


x1 = torch.tensor([[2.], [0.03]])
x2 = torch.tensor([[1.], [0.1]])
calc(x1, x2)

x1 = preprocessing.normalize(x1, axis=0)
x2 = preprocessing.normalize(x2, axis=0)
calc(x1, x2)


def train():
    for i in range(10000):
        # ipdb.set_trace()
        y = w1 * x1 + w2 * x2
        loss = criterion(y, 200 * x1)
        optimizer.zero_grad()
        loss.backward()
        optimizer.step()
        loss_val = loss.item()
        # print(w1.grad, w2.grad)
        if i % 20 == 0:
            print(loss_val)


train()

# (0.57(x-200)-1.04y)^2+(1.8(x-200)-0.2565y)^2

# y = w1 * x1 + w2 * x2
# loss = criterion(y, 20 * x1)

# ((0.5265x + 0.9273y - 200 * 0.5265) ^ 2 + (1.1065x - 1.7421y - 200 * 1.1065) ^ 2) / 2
