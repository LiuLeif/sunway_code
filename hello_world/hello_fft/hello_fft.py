#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2022-09-27 15:02
import matplotlib.pyplot as plt
import numpy as np

sample_rate = 128
ts = 1.0 / sample_rate
t = np.arange(0, 1, ts)

freq = 1.0
x = np.cos(2 * np.pi * freq * t)

freq = 2.0
x += 10 * np.sin(2 * np.pi * freq * t)


def DFT(x):
    N = len(x)
    n = np.arange(N)
    k = n.reshape((N, 1))
    e = np.exp(-2j * np.pi * k * n / N)
    X = np.dot(e, x)
    return X


def IDFT(x):
    N = len(x)
    n = np.arange(N)
    k = n.reshape((N, 1))
    e = np.exp(2j * np.pi * k * n / N)
    X = np.dot(e, x) / N
    return X


X = DFT(x)

print("------")
print(f"1hz cos: {np.real(X[1])/64}")
print(f"2hz sin: {-np.imag(X[2])/64}")
X_1_cos = np.sum(np.cos(-2 * np.pi * t) * x)
X_2_sin = np.sum(np.sin(-2 * 2 * np.pi * t) * x)
print("------")
print(f"1hz cos: {X_1_cos/64}")
print(f"2hz sin: {-X_2_sin/64}")

print("-----")
print("ifft manually:")
for i in range(10):
    x_restored = (
        np.sum(np.cos(i * 2 * np.pi * t) * [np.real(a) for a in X])
        - np.sum(np.sin(i * 2 * np.pi * t) * [np.imag(a) for a in X])
    ) / 128
    print("-----")
    print(x_restored)
    print(x[i])
