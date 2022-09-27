#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2022-09-27 15:02
import matplotlib.pyplot as plt
import numpy as np

sample_rate = 100
ts = 1.0 / sample_rate
t = np.arange(0, 1, ts)

freq = 1.0
x = 3 * np.sin(2 * np.pi * freq * t)

freq = 4.0
x += np.sin(2 * np.pi * freq * t)

# plt.plot(t, x, "r")
# plt.show()


def DFT(x):
    N = len(x)
    n = np.arange(N)
    k = n.reshape((N, 1))
    e = np.exp(-2j * np.pi * k * n / N)
    X = np.dot(e, x)
    return X


X = DFT(x)
N = len(X)
n = np.arange(N)
T = N / sample_rate
freq = n / T

plt.figure(figsize=(8, 6))
plt.stem(freq, abs(X), "b", markerfmt=" ", basefmt="-b")
plt.xlabel("Freq (Hz)")
plt.ylabel("DFT Amplitude |X(freq)|")
plt.show()

# calc X[1] manually
X_1_sin = np.sum(np.sin(2 * np.pi * t) * x)
X_1_cos = np.sum(np.cos(2 * np.pi * t) * x)
