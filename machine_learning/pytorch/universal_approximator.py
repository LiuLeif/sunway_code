#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-08-01 09:41
import numpy as np
import matplotlib.pyplot as plt

epsilon = 0.1


def sigmoid(x):
    return 1. / (1. + np.exp(-x))


def relu(x):
    return np.maximum(0, x)


def bump_relu(h, a, b):
    x = np.linspace(0, 5, 100)
    plt.ylim(0, h * 2)
    plt.plot(x, h / epsilon * (relu(x - a)))

    plt.plot(x, h / epsilon * (relu(x - a - epsilon)))

    left = h / epsilon * (relu(x - a) - relu(x - a - epsilon))
    right = h / epsilon * (relu(x - b) - relu(x - b - epsilon))

    plt.plot(x, left)
    plt.plot(x, left - right)


def bump_sigmoid(h, a, b):
    x = np.linspace(0, 5, 100)
    left = h * sigmoid(1 / epsilon * x - 1 / epsilon * a)
    right = h * sigmoid(1 / epsilon * x - 1 / epsilon * b)

    plt.plot(x, left - right)


# bump_sigmoid(-10, 1, 2)
bump_relu(10, 2, 4)

plt.show()
