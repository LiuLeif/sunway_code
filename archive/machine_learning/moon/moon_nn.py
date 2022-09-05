#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-04-23 23:08
import numpy as np
from PIL import Image
import sys
from sklearn import preprocessing
from matplotlib.colors import ListedColormap
import matplotlib.pyplot as plt

HIDDEN_NODES_NUM = 50
LEARNING_RATE = 0.01
EPOCH = 1000
BATCH_SIZE = 100


def relu(X):
    return np.maximum(X, 0)


def relu_derivative(X):
    return 1. * (X > 0)


def sigmoid(x):
    return 1 / (1 + np.exp(-x))

def cross_entropy(predictions, labels):
    epsilon = 1e-12
    predictions = np.clip(predictions, epsilon, 1 - epsilon)
    N = predictions.shape[0]

    ce = 0 - np.sum(np.sum(labels * np.log(predictions))) / N
    return ce


def cost_function(X, Y, W1, B1, W2, B2):
    # X:   m   *  2
    # Y:   m   *  1
    # W1:  2   *  HIDDEN_NODES_NUM
    # B1:  1   *  HIDDEN_NODES_NUM
    # W2:  HIDDEN_NODES_NUM   *  1
    # B2:  1   *  1

    m = len(X)
    J = 0.
    dw1 = np.zeros_like(W1)
    db1 = np.zeros_like(B1)
    dw2 = np.zeros_like(W2)
    db2 = np.zeros_like(B2)
    # for i in range(m):
    #     x = X[i, :].reshape(1, 2)
    #     y = Y[i, :].reshape(1, 1)
    #     z1 = np.matmul(x, W1) + B1
    #     a1 = relu(z1)
    #     z2 = np.matmul(a1, W2) + B2
    #     y_hat = sigmoid(z2)
    #     loss = cross_entropy(
    #            np.concatenate((y_hat, 1 - y_hat)), np.concatenate((y, 1 - y)))
    #     J += loss
    #     # bp
    #     delta3 = y_hat - y
    #     dw2 += np.matmul(np.transpose(a1), delta3)
    #     db2 += delta3
    #     # 1*HIDDEN_NODES_NUM
    #     delta2 = np.matmul(delta3, np.transpose(W2))
    #     delta2 = delta2 * relu_derivative(a1)

    #     # dw1: 784 * HIDDEN_NODES_NUM
    #     dw1 += np.matmul(np.transpose(x), delta2)
    #     # db1: 1 * HIDDEN_NODES_NUM
    #     db1 += delta2
    # J /= m
    # dw1 /= m
    # db1 /= m
    # dw2 /= m
    # db2 /= m

    z1 = np.matmul(X, W1) + B1
    a1 = relu(z1)
    z2 = np.matmul(a1, W2) + B2
    y_hat = sigmoid(z2)
    J = cross_entropy(
        np.concatenate((y_hat, 1 - y_hat)), np.concatenate((Y, 1 - Y)))

    # bp
    delta3 = y_hat - Y
    dw2 = np.matmul(np.transpose(a1), delta3) / m
    db2 = delta3.mean(axis=0)

    # 1*HIDDEN_NODES_NUM
    delta2 = np.matmul(delta3, np.transpose(W2))
    delta2 = delta2 * relu_derivative(a1)

    # dw1: 784 * HIDDEN_NODES_NUM
    dw1 = np.matmul(np.transpose(X), delta2) / m
    # db1: 1 * HIDDEN_NODES_NUM
    db1 = delta2.mean(axis=0)
    return J, dw1, db1, dw2, db2


from os import listdir
from os.path import isfile, join

from sklearn.datasets import make_moons


def get_training_set():
    X, Y = make_moons(n_samples=1000, noise=0.2)
    Y = Y.reshape(1000, 1)
    X = preprocessing.scale(X)
    return X, Y


def gradient_decent(X, Y, W1, B1, W2, B2):
    alpha = LEARNING_RATE
    for epoch in range(EPOCH):
        batch = len(X) // BATCH_SIZE
        total_loss = 0
        for X_batch, Y_batch in zip(
                np.split(X[:batch * BATCH_SIZE], batch),
                np.split(Y[:batch * BATCH_SIZE], batch)):
            cost, dw1, db1, dw2, db2 = cost_function(X_batch, Y_batch, W1, B1,
                                                     W2, B2)
            # print("training: #", epoch, cost)
            total_loss += cost
            W1 = W1 - alpha * dw1
            B1 = B1 - alpha * db1
            W2 = W2 - alpha * dw2
            B2 = B2 - alpha * db2
        if epoch % (EPOCH // 50) == 0:
            print("training: #", epoch, total_loss / batch)
    return W1, B1, W2, B2


def draw_decision_boundary(W1, B1, W2, B2):
    xx, yy = np.meshgrid(np.arange(-4, 4, 0.02), np.arange(-4, 4, 0.02))
    X = np.c_[xx.ravel(), yy.ravel()]

    z1 = np.dot(X, W1) + B1
    a1 = relu(z1)
    z2 = np.dot(a1, W2) + B2
    a2 = sigmoid(z2)
    y_hat = (a2 > 0.5)[:, 0]
    y_hat = y_hat.reshape(xx.shape)
    cm = ListedColormap(['#FF0000', '#0000FF'])
    plt.contour(xx, yy, y_hat, cmap=cm)


def train():
    global W1, B1, W2, B2
    X, Y = get_training_set()
    Z = np.concatenate((X, Y), axis=1)
    np.random.shuffle(Z)
    X = Z[:, :2]
    Y = Z[:, 2:]
    W1 = np.random.randn(2, HIDDEN_NODES_NUM)
    B1 = np.random.randn(1, HIDDEN_NODES_NUM)
    W2 = np.random.randn(HIDDEN_NODES_NUM, 1)
    B2 = np.random.randn(1, 1)
    W1, B1, W2, B2 = gradient_decent(X, Y, W1, B1, W2, B2)


W1, B1, W2, B2 = 0, 0, 0, 0


def predict():
    global W1, B1, W2, B2
    X, Y = get_training_set()
    cm = ListedColormap(['#FF0000', '#0000FF'])
    plt.scatter(x=X[:, 0], y=X[:, 1], c=Y[:, 0], cmap=cm)
    draw_decision_boundary(W1, B1, W2, B2)
    plt.show()


train()
predict()
