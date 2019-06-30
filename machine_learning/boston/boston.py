#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-04-23 23:08
from sklearn import preprocessing
from sklearn.datasets import load_boston
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.colors import ListedColormap
from PIL import Image
import sys

EPOCH = 1000
LEARNING_RATE = 0.01


def MSE(A, B):
    return np.square(np.subtract(A, B)).mean()


def get_training_set():
    X, Y = load_boston(True)
    X = preprocessing.scale(X)
    Y = Y.reshape(-1, 1)
    [m, features] = X.shape
    # Z = np.concatenate((X, Y), axis=1)
    # np.random.shuffle(Z)
    # X = Z[:, :features]
    # Y = Z[:, features:]
    global X_train, Y_train, X_test, Y_test
    offset = int(0.8 * m)
    X_train, Y_train = X[:offset], Y[:offset]
    X_test, Y_test = X[offset:], Y[offset:]


def cost_function(X, Y, W, B):
    m = len(X)
    J = 0.
    dw = np.zeros_like(W)
    db = np.zeros_like(B)
    # for i in range(m):
    #     x = X[i, :].reshape(1, -1)
    #     y = Y[i, :].reshape(1, 1)
    #     y_hat = np.matmul(x, W) + B
    #     # print(y_hat[0], y[0], MSE(y_hat[0], y[0]))
    #     loss = MSE(y_hat, y)
    #     J += loss
    #     dw += np.matmul(np.transpose(x), y_hat - y)
    #     db += y_hat - y
    # J /= m
    # dw /= m
    # db /= m
    y_hat = np.matmul(X, W) + B
    J = MSE(y_hat, Y)
    dw = np.matmul(np.transpose(X), y_hat - Y) / m
    db = (y_hat - Y).mean(axis=0)
    return J, dw, db


def gradient_decent(X, Y, W, B):
    alpha = LEARNING_RATE
    for epoch in range(EPOCH):
        cost, dw, db = cost_function(X, Y, W, B)
        W = W - alpha * dw
        B = B - alpha * db
        if epoch % (EPOCH // 30) == 0:
            print("training: #", epoch, cost)
    return W, B


X_test, Y_test, X_train, Y_train = 0, 0, 0, 0
W, B = 0, 0


def train():
    global X_train, Y_train, W, B
    W = np.random.randn(X_train.shape[1], 1)
    B = np.random.randn(1, 1)
    W, B = gradient_decent(X_train, Y_train, W, B)


def predict():
    global W, B
    predicted = np.matmul(X_test, W) + B
    print(MSE(predicted, Y_test))


get_training_set()
train()
predict()
