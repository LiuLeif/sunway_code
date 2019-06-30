#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-04-23 23:08
import numpy as np
import sys
from sklearn import preprocessing
from sklearn.datasets import load_digits

HIDDEN_NODES_NUM = 30
LEARNING_RATE = 0.01
EPOCH = 100
BATCH_SIZE = 2
REGULARIZATION_FACTOR = 0.01
USE_RELU = True


def relu(X):
    return np.maximum(X, 0)


def relu_derivative(X):
    return 1. * (X > 0)


def sigmoid(x):
    return 1 / (1 + np.exp(-x))


def sigmoid_derivative(X):
    return X * (1 - X)


def softmax(z):
    assert len(z.shape) == 2
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


def cost_function(X, Y, W1, B1, W2, B2):
    m = len(X)
    assert (len(X) == len(Y))

    J = 0.

    dw1 = np.zeros_like(W1)
    db1 = np.zeros_like(B1)

    dw2 = np.zeros_like(W2)
    db2 = np.zeros_like(B2)
    # normal
    # for i in range(m):
    #     x = X[i, :].reshape(1, -1)
    #     y = Y[i, :].reshape(1, -1)
    #     z1 = np.matmul(x, W1) + B1
    #     if USE_RELU:
    #         a1 = relu(z1)
    #     else:
    #         a1 = sigmoid(z1)
    #     z2 = np.matmul(a1, W2) + B2
    #     y_hat = softmax(z2)
    #     loss = cross_entropy(y_hat, y)
    #     J += loss
    #     # bp
    #     delta3 = y_hat - y
    #     dw2 += np.matmul(np.transpose(a1), delta3)
    #     db2 += delta3
    #     delta2 = np.matmul(delta3, np.transpose(W2))
    #     if USE_RELU:
    #         delta2 = delta2 * relu_derivative(a1)
    #     else:
    #         delta2 = delta2 * sigmoid_derivative(a1)
    #     dw1 += np.matmul(np.transpose(x), delta2)
    #     db1 += delta2
    # J += REGULARIZATION_FACTOR * (
    #     np.sum(np.sum(np.square(W1))) + np.sum(np.sum(np.square(W2)))) / 2
    # dw1 += REGULARIZATION_FACTOR * W1
    # dw2 += REGULARIZATION_FACTOR * W2
    # J /= m
    # dw1 /= m
    # db1 /= m
    # dw2 /= m
    # db2 /= m

    # vectorization
    z1 = np.matmul(X, W1) + B1
    if USE_RELU:
        a1 = relu(z1)
    else:
        a1 = sigmoid(z1)
    z2 = np.matmul(a1, W2) + B2
    y_hat = softmax(z2)
    J = cross_entropy(y_hat, Y)

    delta3 = y_hat - Y
    dw2 = np.matmul(np.transpose(a1), delta3) / m
    db2 = delta3.mean(axis=0)

    delta2 = np.matmul(delta3, np.transpose(W2))
    if USE_RELU:
        delta2 = delta2 * relu_derivative(a1)
    else:
        delta2 = delta2 * sigmoid_derivative(a1)

    dw1 = np.matmul(np.transpose(X), delta2) / m
    db1 = delta2.mean(axis=0)

    # regulation
    J += REGULARIZATION_FACTOR * (
        np.sum(np.sum(np.square(W1))) + np.sum(np.sum(np.square(W2)))) / (
            2 * m)
    dw1 += REGULARIZATION_FACTOR * W1 / m
    dw2 += REGULARIZATION_FACTOR * W2 / m

    return J, dw1, db1, dw2, db2


X_train, Y_train, X_test, Y_test = 0, 0, 0, 0
W1, B1, W2, B2 = 0, 0, 0, 0


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
    print(X.shape)
    print(Y_train.shape)
    print(Y_test.shape)


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
            total_loss += cost
            W1 = W1 - alpha * dw1
            B1 = B1 - alpha * db1
            W2 = W2 - alpha * dw2
        if epoch % (EPOCH // 30) == 0:
            print("training: #", epoch, total_loss / batch)

    return W1, B1, W2, B2


def predict():
    global X_test, Y_test, W1, B1, W2, B2

    wrong = 0
    correct = 0
    for i in range(len(X_test)):
        x = X_test[i, :].reshape(1, -1)
        y = Y_test[i, :].reshape(1, -1)
        z1 = np.matmul(x, W1) + B1
        if USE_RELU:
            a1 = relu(z1)
        else:
            a1 = sigmoid(z1)
        z2 = np.matmul(a1, W2) + B2
        a2 = softmax(z2)

        if np.argmax(a2) == np.argmax(y):
            correct += 1
        else:
            wrong += 1
    print("correct: %d, wrong: %d, accuracy: %f" % (correct, wrong, correct /
                                                    (correct + wrong)))


def train():
    global W1, B2, W2, B2, X_train, Y_train
    W1 = np.random.randn(64, HIDDEN_NODES_NUM)
    B1 = np.random.randn(1, HIDDEN_NODES_NUM)

    W2 = np.random.randn(HIDDEN_NODES_NUM, 10)
    B2 = np.random.randn(1, 10)

    W1, B1, W2, B2 = gradient_decent(X_train, Y_train, W1, B1, W2, B2)


def check_gradient():
    global X_train, Y_train
    W1 = np.random.randn(64, HIDDEN_NODES_NUM)
    B1 = np.random.randn(1, HIDDEN_NODES_NUM)

    W2 = np.random.randn(HIDDEN_NODES_NUM, 10)
    B2 = np.random.randn(1, 10)

    K1 = np.random.randn(3, 9)
    KB1 = np.random.randn(3, 1)

    cost, dw1, db1, dw2, db2 = cost_function(X_train, Y_train, W1, B1, W2, B2)

    W1[0, 0] += 1e-3
    cost2, _, _, _, _ = cost_function(X_train, Y_train, W1, B1, W2, B2)
    W1[0, 0] -= 1e-3
    print((cost2 - cost) / 1e-3, dw1[0, 0])


get_training_set()
check_gradient()
train()
predict()
