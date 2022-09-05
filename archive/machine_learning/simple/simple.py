#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-04-23 23:08
from sklearn import preprocessing
import numpy as np
import matplotlib.pyplot as plt

EPOCH = 1000
LEARNING_RATE = 0.01


def MSE(A, B):
    return np.square(np.subtract(A, B)).mean()


def get_training_set():
    global X, Y
    data = np.loadtxt("data.txt", delimiter=",")
    X = data[:, 0].reshape(-1, 1)
    X = preprocessing.scale(X)
    Y = data[:, 1].reshape(-1, 1)
    Y = preprocessing.scale(Y)


def cost_function(X, Y, W, B):
    m = len(X)
    J = 0.
    dw = np.zeros_like(W)
    db = np.zeros_like(B)

    # for i in range(m):
    #     x = X[i, :].reshape(1, -1)
    #     y = Y[i, :].reshape(1, 1)
    #     y_hat = np.matmul(x, W) + B
    #     loss = MSE(y_hat, y)
    #     J += loss
    #     dw += np.matmul(np.transpose(x), y_hat - y)
    #     db += y_hat - y
    # J /= m
    # dw /= m
    # db /= m

    # vectorization
    y_hat = np.matmul(X, W) + B
    J = MSE(y_hat, Y)
    dw = np.matmul(np.transpose(X), y_hat - Y) / m
    db = (y_hat - Y).mean()
    return J, dw, db


def gradient_decent(X, Y, W, B):
    alpha = LEARNING_RATE
    for epoch in range(EPOCH):
        cost, dw, db = cost_function(X, Y, W, B)
        W = W - alpha * dw
        B = B - alpha * db
        print("training: #", epoch, cost)
    return W, B


X, Y = 0, 0
W, B = 0, 0


def train():
    global X, Y, W, B
    W = np.random.randn(X.shape[1], 1)
    B = np.random.randn(1, 1)
    W, B = gradient_decent(X, Y, W, B)


def predict():
    global W, B, X, Y
    plt.scatter(x=X[:, 0], y=Y[:, 0])
    plt.plot(X, np.matmul(X, W) + B, "r")
    plt.show()


get_training_set()
train()
predict()
