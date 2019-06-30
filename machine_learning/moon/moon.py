#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-04-23 23:08
from sklearn import preprocessing
from sklearn.datasets import make_moons
from sklearn.datasets import make_circles
import numpy as np
import matplotlib.pyplot as plt
from matplotlib.colors import ListedColormap
from PIL import Image
import sys

EPOCH = 1000
LEARNING_RATE = 0.01
BATCH_SIZE = 10
FEATURE_SIZE = 0
POLY_FEATURES = 3


def sigmoid(x):
    return 1 / (1 + np.exp(-x))


def cross_entropy(predictions, labels):
    epsilon = 1e-12
    predictions = np.clip(predictions, epsilon, 1 - epsilon)
    N = predictions.shape[0]
    ce = 0 - np.sum(np.sum(labels * np.log(predictions))) / N
    return ce


def draw_decision_boundary(W, B):
    xx, yy = np.meshgrid(
        np.arange(-1.5, 2.5, 0.02), np.arange(-1.5, 2.5, 0.02))
    X = np.c_[xx.ravel(), yy.ravel()]

    poly = preprocessing.PolynomialFeatures(POLY_FEATURES, include_bias=False)
    X = poly.fit_transform(X)

    Y = ((sigmoid(np.matmul(X, W) + B)) > 0.5)[:, 0]
    Y = Y.reshape(xx.shape)
    cm = ListedColormap(['#FF0000', '#0000FF'])
    plt.contour(xx, yy, Y, cmap=cm)


def get_training_set():
    X, Y = make_moons(n_samples=1000, noise=0.2)
    # X, Y = make_circles(n_samples=1000, noise=0.2, factor=0.5)
    Y = Y.reshape(1000, 1)
    poly = preprocessing.PolynomialFeatures(POLY_FEATURES, include_bias=False)
    X = poly.fit_transform(X, Y)
    global FEATURE_SIZE
    FEATURE_SIZE = X.shape[1]
    return X, Y


def cost_function(X, Y, W, B):
    # X: m * 2
    # Y: m * 1
    # W: 2 * 1
    # B: 1 * 1

    m = len(X)
    assert (len(X) == len(Y))

    J = 0.

    dw = np.zeros_like(W)
    db = np.zeros_like(B)

    # for i in range(m):
    #     x = X[i, :].reshape(1, FEATURE_SIZE)
    #     y = Y[i, :].reshape(1, 1)
    #     f = np.matmul(x, W) + B
    #     y_hat = sigmoid(f)
    #     loss = cross_entropy(
    #         np.concatenate((y_hat, 1 - y_hat)), np.concatenate((y, 1 - y)))
    #     J += loss

    #     dw += np.matmul(np.transpose(x), y_hat - y)
    #     db += y_hat - y

    # # regulation
    # # J += 0.1 * sum(sum(np.square(W))) / 2
    # # dw += 0.1 * W
    # J /= m
    # dw /= m
    # db /= m

    y_hat = sigmoid(np.matmul(X, W) + B)
    J = cross_entropy(
        np.concatenate((y_hat, 1 - y_hat)), np.concatenate((Y, 1 - Y)))
    dw = np.matmul(X.T, y_hat - Y) / m
    db = (y_hat - Y).mean(axis=0)
    return J, dw, db


def gradient_decent(X, Y, W, B):
    alpha = LEARNING_RATE
    for epoch in range(EPOCH):
        batch = len(X) // BATCH_SIZE
        total_loss = 0
        for X_batch, Y_batch in zip(
                np.split(X[:batch * BATCH_SIZE], batch),
                np.split(Y[:batch * BATCH_SIZE], batch)):
            cost, dw, db = cost_function(X_batch, Y_batch, W, B)
            total_loss += cost
            W = W - alpha * dw
            B = B - alpha * db
        if epoch % (EPOCH // 30) == 0:
            print("training: #", epoch, total_loss / batch)

    return W, B


def train():
    global W, B
    X, Y = get_training_set()

    Z = np.concatenate((X, Y), axis=1)
    np.random.shuffle(Z)
    X = Z[:, :FEATURE_SIZE]
    Y = Z[:, FEATURE_SIZE:]

    W = np.random.randn(FEATURE_SIZE, 1)
    B = np.random.randn(1, 1)

    W, B = gradient_decent(X, Y, W, B)


W, B = 0, 0


def predict():
    global W, B
    X, Y = get_training_set()
    cm = ListedColormap(['#FF0000', '#0000FF'])
    # plt.plot(X[:, 0], X[:, 1], 'bo', label='Real data')
    y_hat = ((sigmoid(np.matmul(X, W) + B)) > 0.5)[:, 0]
    plt.scatter(x=X[:, 0], y=X[:, 1], c=Y[:, 0], cmap=cm)
    # plt.scatter(x=X[:, 0], y=X[:, 1], c=y_hat, cmap=cm)
    draw_decision_boundary(W, B)
    plt.show()


train()
predict()
