#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-04-23 23:08
import numpy as np
from PIL import Image
import sys
from sklearn import preprocessing
from sklearn.datasets import load_digits

EPOCH = 100
LEARNING_RATE = 0.01
BATCH_SIZE = 2
REGULARIZATION_FACTOR = 0.01


def softmax(z):
    s = np.max(z, axis=1)
    s = s[:, np.newaxis]  # necessary step to do broadcasting
    e_x = np.exp(z - s)
    div = np.sum(e_x, axis=1)
    div = div[:, np.newaxis]  # dito
    return e_x / div


def one_hot(Y, C):
    Y = np.eye(C)[Y.reshape(-1)]
    return Y


def cross_entropy(predictions, labels):
    epsilon = 1e-12
    predictions = np.clip(predictions, epsilon, 1 - epsilon)
    N = predictions.shape[0]

    ce = 0 - np.sum(np.sum(labels * np.log(predictions))) / N
    return ce


def cost_function(X, Y, W, B):
    m = len(X)
    assert (len(X) == len(Y))

    J = 0.

    dw = np.zeros_like(W)
    db = np.zeros_like(B)

    # for i in range(m):
    #     x = X[i, :].reshape(1, 64)
    #     y = Y[i, :].reshape(1, 10)
    #     f = np.matmul(x, W) + B
    #     assert (f.shape == (1, 10))
    #     y_hat = softmax(f)
    #     assert (y.shape == y_hat.shape)
    #     loss = cross_entropy(y_hat, y)
    #     J += loss

    #     dw += np.matmul(np.transpose(x), y_hat - y)
    #     db += y_hat - y

    # J /= m
    # dw /= m
    # db /= m

    f = np.matmul(X, W) + B
    y_hat = softmax(f)
    J = cross_entropy(y_hat, Y)
    dw = np.matmul(np.transpose(X), y_hat - Y) / m
    db = (y_hat - Y).mean(axis=0)

    # regularation
    J += REGULARIZATION_FACTOR * sum(sum(np.square(W))) / (2 * m)
    dw += REGULARIZATION_FACTOR * W / m
    return J, dw, db


X_train, Y_train, X_test, Y_test = 0, 0, 0, 0
W, B = 0, 0


def get_training_set():
    X, Y = load_digits(10, True)
    Y = Y.reshape(-1, 1)
    Y = one_hot(Y, 10)
    [m, features] = X.shape
    Z = np.concatenate((X, Y), axis=1)
    np.random.shuffle(Z)
    X = Z[:, :features]
    Y = Z[:, features:]
    global X_train, Y_train, X_test, Y_test
    offset = int(0.8 * m)
    X_train, Y_train = X[:offset], Y[:offset]
    X_test, Y_test = X[offset:], Y[offset:]


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


def predict():
    global X_test, Y_test, W, B

    wrong = 0
    correct = 0
    for i in range(len(X_test)):
        x = X_test[i, :].reshape(1, -1)
        y = Y_test[i, :].reshape(1, -1)
        c = np.matmul(x, W) + B
        y_hat = softmax(c)

        if np.argmax(y_hat) == np.argmax(y):
            correct += 1
        else:
            wrong += 1
    print("correct: %d, wrong: %d, accuracy: %f" % (correct, wrong, correct /
                                                    (correct + wrong)))


def train():
    global W, B, X_train, Y_train
    W = np.random.randn(64, 10)
    B = np.random.randn(1, 10)
    W, B = gradient_decent(X_train, Y_train, W, B)


get_training_set()
train()
predict()
