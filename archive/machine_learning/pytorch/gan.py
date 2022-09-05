#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-06-13 15:12
import numpy as np
import torch
from torch import nn
from torch import optim
from tensorboardX import SummaryWriter

LEARNING_RATE = 2e-4

def decorate_with_diffs(data, exponent):
    mean = torch.mean(data.data, 1, keepdim=True)
    mean_broadcast = torch.mul(torch.ones(data.size()), mean.tolist()[0][0])
    diffs = torch.pow(data - mean_broadcast, exponent)
    return torch.cat([data, diffs], 1)

(name, preprocess, d_input_func) = ("Data and variances", lambda data: decorate_with_diffs(data, 2.0), lambda x: x * 2)


def get_distribution(mu, sigma):
    return lambda n: torch.from_numpy(np.random.normal(mu, sigma, (1, n))).float()


real_data_sampler = get_distribution(10, 2)


def get_generator_input():
    return lambda m, n: torch.rand(m, n).float()


fake_data_sampler = get_generator_input()

HIDDEN_SIZE = 50

g_input_size = 1
g_output_size = 1

d_input_size = 100
d_output_size = 1

generator = nn.Sequential(
    nn.Linear(g_input_size, HIDDEN_SIZE),
    nn.ELU(),
    nn.Linear(HIDDEN_SIZE, HIDDEN_SIZE),
    nn.Sigmoid(),
    nn.Linear(HIDDEN_SIZE, g_output_size),
)

discriminator = nn.Sequential(
    nn.Linear(d_input_func(d_input_size), HIDDEN_SIZE), nn.ELU(),
    nn.Linear(HIDDEN_SIZE, HIDDEN_SIZE), nn.ELU(),
    nn.Linear(HIDDEN_SIZE, d_output_size), nn.Sigmoid())

criterion = nn.BCELoss()
g_optimizer = optim.Adam(generator.parameters(), lr=LEARNING_RATE)
d_optimizer = optim.Adam(discriminator.parameters(), lr=LEARNING_RATE)

writer = SummaryWriter()

for epoch in range(100000):
    # train discriminator
    discriminator.zero_grad()
    d_real_data = real_data_sampler(d_input_size)
    d_real_decision = discriminator(preprocess(d_real_data))
    d_real_loss = criterion(d_real_decision, torch.ones(1))
    d_real_loss.backward()

    x = fake_data_sampler(d_input_size, g_input_size)
    d_fake_data = generator(x).detach()
    d_fake_decision = discriminator(preprocess(d_fake_data.t()))
    d_fake_loss = criterion(d_fake_decision, torch.zeros(1))
    d_fake_loss.backward()
    d_optimizer.step()

    # train generator
    generator.zero_grad()
    d_fake_data = generator(fake_data_sampler(d_input_size, g_input_size))

    d_fake_decision = discriminator(preprocess(d_fake_data.t()))
    g_loss = criterion(d_fake_decision, torch.ones(1))
    g_loss.backward()

    g_optimizer.step()

    # writer.add_scalar("d_fake_decision", d_fake_decision.item(), epoch)
    if epoch % 200 == 0:
        print(
            "training: #%d g_loss: %f, d_real: %f, d_fake: %f, mean: %f, std: %f"
            % (epoch, g_loss.item(), d_real_loss.item(), d_fake_loss.item(),
               d_fake_data.mean().item(), d_fake_data.std().item()))
        
writer.close()

import matplotlib.pyplot as plt

def plot():
    x = fake_data_sampler(10000, g_input_size)
    d_fake_data = generator(x).detach()
    # d_fake_data = real_data_sampler(10000)
    x = d_fake_data.numpy()
    x = x.reshape(-1)
    plt.hist(x, bins=1000)
    plt.show()
