#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-07-20 22:17
import numpy as np
import matplotlib.pyplot as plt
import torch
import torch.nn.functional as F
from tensorboardX import SummaryWriter

plt.style.use("classic")

alpha = 0.15

# criterion = torch.nn.MSELoss(reduce=False)
criterion = torch.nn.BCELoss(reduce=False)


def update_gd(x, y, step):
    xx = torch.tensor(x, requires_grad=True)
    yy = torch.tensor(y, requires_grad=True)
    zz = F.sigmoid(xx + yy)
    loss = criterion(zz, torch.tensor(0.))
    writer.add_scalar("loss", loss, step)
    loss.backward()
    print(zz.item(), xx.grad.item())
    return x - alpha * xx.grad, y - alpha * yy.grad


plt.axis("equal")

x = np.linspace(-10, 10, 50)
y = np.linspace(-10, 10, 50)
xx, yy = np.meshgrid(x, y)

z = F.sigmoid(torch.tensor(xx + yy).unsqueeze(0))
loss = criterion(z, torch.zeros_like(z)).squeeze(0).numpy()
CS = plt.contour(xx, yy, loss, 20)
plt.clabel(CS, inline=1, fontsize=10)

writer = SummaryWriter("/tmp/runs/slow")
x, y = 2., 2.
for i in range(200):
    plt.scatter(x, y)
    x2, y2 = update_gd(x, y, i)
    plt.plot([x, x2], [y, y2])
    x, y = x2, y2
writer.close()

print(x, y)
# writer = SummaryWriter("/tmp/runs/fast")
# x, y = 2., 0.
# for i in range(200):
#     plt.scatter(x, y)
#     x2, y2 = update_gd(x, y, i)
#     plt.plot([x, x2], [y, y2])
#     x, y = x2, y2
# writer.close()

plt.show()
print(x, y)
