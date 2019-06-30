#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-04-23 23:08
import numpy as np
import sys
from sklearn import preprocessing
from sklearn.datasets import load_digits
from im2col import *
import matplotlib.pyplot as plt

HIDDEN_NODES_NUM = 100
LEARNING_RATE = 0.1
EPOCH = 200
BATCH_SIZE = 32
REGULARIZATION_FACTOR = 0.001
USE_RELU = False


def max_pooling_backward(X, da0):
    # X: 1 * 3 * 8 * 8
    # return: 3*4*4
    X_reshaped = X.reshape(3, 1, 8, 8)
    X_col = im2col_indices(X_reshaped, 2, 2, 0, 2)
    max_idx = np.argmax(X_col, axis=0)
    dX_col = np.zeros_like(X_col)
    dout_flat = da0.transpose(2, 3, 0, 1).ravel()
    dX_col[max_idx, range(max_idx.size)] = dout_flat
    dX = col2im_indices(dX_col, (3, 1, 8, 8), 2, 2, 0, 2)
    dX = dX.reshape(X.shape)
    return dX


def max_pooling_forward(X):
    # X: 1 * 3 * 8 * 8
    # return: 3*4*4
    X_reshaped = X.reshape(3, 1, 8, 8)
    X_col = im2col_indices(X_reshaped, 2, 2, 0, 2)
    max_idx = np.argmax(X_col, axis=0)
    out = X_col[max_idx, range(max_idx.size)]
    out = out.reshape(4, 4, 1, 3)
    out = out.transpose(2, 3, 0, 1)
    return out


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


def cost_function(X, Y, W1, B1, W2, B2, K1, KB1):
    m = len(X)
    assert (len(X) == len(Y))

    J = 0.

    dw1 = np.zeros_like(W1)
    db1 = np.zeros_like(B1)

    dw2 = np.zeros_like(W2)
    db2 = np.zeros_like(B2)

    dk1 = np.zeros_like(K1)
    dkb1 = np.zeros_like(KB1)

    # normal
    for i in range(m):
        # forward
        x = X[i, :].reshape(1, 1, 8, -1)
        y = Y[i, :].reshape(1, -1)

        col = im2col_indices(x, 3, 3, 1, 1)
        c00 = np.dot(K1, col) + KB1
        # c00 = c00.flatten().reshape(1, -1)
        if USE_RELU:
            a00 = relu(c00)
        else:
            a00 = sigmoid(c00)
        # max_pooling layer
        # a00: 1,192 = 3 * 8 * 8
        a00 = a00.reshape(1, 3, 8, 8)
        # a0: 3 * 4 * 4
        a0 = max_pooling_forward(a00)
        a0 = a0.reshape(1, -1)
        # fc layer
        z1 = np.matmul(a0, W1) + B1
        if USE_RELU:
            a1 = relu(z1)
        else:
            a1 = sigmoid(z1)
        z2 = np.matmul(a1, W2) + B2
        y_hat = softmax(z2)
        loss = cross_entropy(y_hat, y)
        J += loss
        # backward
        # w2
        da2 = y_hat - y
        dw2 += np.matmul(np.transpose(a1), da2)
        db2 += da2
        da1 = np.matmul(da2, np.transpose(W2))
        if USE_RELU:
            da1 = da1 * relu_derivative(a1)
        else:
            da1 = da1 * sigmoid_derivative(a1)
        # w1
        dw1 += np.matmul(np.transpose(a0), da1)
        db1 += da1
        da0 = np.matmul(da1, np.transpose(W1))
        da0 = da0.reshape(1, 3, 4, 4)
        # max pooling
        # convert 3* 16 to 3*64
        da00 = max_pooling_backward(a00, da0)
        a00 = a00.reshape(3, -1)
        da00 = da00.reshape(3, -1)
        if USE_RELU:
            da00 = da00 * relu_derivative(a00)
        else:
            da00 = da00 * sigmoid_derivative(a00)
        dk1 += np.matmul(da00, np.transpose(col))
        dkb1 += np.sum(da00, axis=1).reshape(3, -1)

    # regularization
    # J += REGULARIZATION_FACTOR * (np.sum(np.sum(np.square(W1))) + np.sum(
    #     np.sum(np.square(W2))) + np.sum(np.sum(np.square(K1)))) / 2
    # dw1 += REGULARIZATION_FACTOR * W1
    # dw2 += REGULARIZATION_FACTOR * W2
    # dk1 += REGULARIZATION_FACTOR * K1

    J /= m
    dw1 /= m
    db1 /= m
    dw2 /= m
    db2 /= m
    dk1 /= m
    dkb1 /= m

    # vectorization
    # z1 = np.matmul(X, W1) + B1
    # if USE_RELU:
    #     a1 = relu(z1)
    # else:
    #     a1 = sigmoid(z1)
    # z2 = np.matmul(a1, W2) + B2
    # y_hat = softmax(z2)
    # J = cross_entropy(y_hat, Y)

    # delta3 = y_hat - Y
    # dw2 = np.matmul(np.transpose(a1), delta3) / m
    # db2 = delta3.mean(axis=0)

    # delta2 = np.matmul(delta3, np.transpose(W2))
    # if USE_RELU:
    #     delta2 = delta2 * relu_derivative(a1)
    # else:
    #     delta2 = delta2 * sigmoid_derivative(a1)

    # dw1 = np.matmul(np.transpose(X), delta2) / m
    # db1 = delta2.mean(axis=0)

    # regulation
    # J += REGULARIZATION_FACTOR * (
    #     np.sum(np.sum(np.square(W1))) + np.sum(np.sum(np.square(W2)))) / (
    #         2 * m)
    # dw1 += REGULARIZATION_FACTOR * W1 / m
    # dw2 += REGULARIZATION_FACTOR * W2 / m

    return J, dw1, db1, dw2, db2, dk1, dkb1


X_train, Y_train, X_test, Y_test = 0, 0, 0, 0
W1, B1, W2, B2 = 0, 0, 0, 0
K1, K2, K3 = 0, 0, 0


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


def gradient_decent(X, Y, W1, B1, W2, B2, K1, KB1):
    alpha = LEARNING_RATE
    for epoch in range(EPOCH):
        batch = len(X) // BATCH_SIZE
        total_loss = 0
        for X_batch, Y_batch in zip(
                np.split(X[:batch * BATCH_SIZE], batch),
                np.split(Y[:batch * BATCH_SIZE], batch)):
            cost, dw1, db1, dw2, db2, dk1, dkb1 = cost_function(
                X_batch, Y_batch, W1, B1, W2, B2, K1, KB1)
            total_loss += cost
            W1 = W1 - alpha * dw1
            B1 = B1 - alpha * db1
            W2 = W2 - alpha * dw2
            K1 = K1 - alpha * dk1
            KB1 = KB1 - alpha * dkb1
        if epoch % (EPOCH // 30) == 0:
            print("training: #", epoch, total_loss / batch)

    return W1, B1, W2, B2, K1, KB1


def predict():
    global X_test, Y_test, W1, B1, W2, B2, K1, KB1

    wrong = 0
    correct = 0
    for i in range(len(X_test)):
        x = X_test[i, :].reshape(1, 1, 8, -1)
        y = Y_test[i, :].reshape(1, -1)

        col = im2col_indices(x, 3, 3, 1, 1)
        c00 = np.dot(K1, col) + KB1
        c00 = c00.flatten().reshape(1, -1)
        if USE_RELU:
            a00 = relu(c00)
        else:
            a00 = sigmoid(c00)

        a0 = max_pooling_forward(a00)
        a0 = a0.reshape(1, -1)
        z1 = np.matmul(a0, W1) + B1
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
    global K1, KB1
    W1 = np.random.randn(48, HIDDEN_NODES_NUM)
    B1 = np.random.randn(1, HIDDEN_NODES_NUM)

    W2 = np.random.randn(HIDDEN_NODES_NUM, 10)
    B2 = np.random.randn(1, 10)

    K1 = np.random.randn(3, 9)
    KB1 = np.random.randn(3, 1)

    W1, B1, W2, B2, K1, KB1 = gradient_decent(X_train, Y_train, W1, B1, W2, B2,
                                              K1, KB1)


def check_gradient():
    global X_train, Y_train
    W1 = np.random.randn(48, HIDDEN_NODES_NUM)
    B1 = np.random.randn(1, HIDDEN_NODES_NUM)

    W2 = np.random.randn(HIDDEN_NODES_NUM, 10)
    B2 = np.random.randn(1, 10)

    K1 = np.random.randn(3, 9)
    KB1 = np.random.randn(3, 1)

    cost, dw1, db1, dw2, db2, dk1, dkb1 = cost_function(
        X_train, Y_train, W1, B1, W2, B2, K1, KB1)

    K1[0, 0] += 1e-4
    cost2, _, _, _, _, _, _ = cost_function(X_train, Y_train, W1, B1, W2, B2,
                                            K1, KB1)
    K1[0, 0] -= 1e-4
    print((cost2 - cost) / 1e-4, dk1[0, 0])

    KB1[0, 0] += 1e-4
    cost2, _, _, _, _, _, _ = cost_function(X_train, Y_train, W1, B1, W2, B2,
                                            K1, KB1)
    KB1[0, 0] -= 1e-4
    print((cost2 - cost) / 1e-4, dkb1[0, 0])

    W1[0, 0] += 1e-4
    cost2, _, _, _, _, _, _ = cost_function(X_train, Y_train, W1, B1, W2, B2,
                                            K1, KB1)
    W1[0, 0] -= 1e-4
    print((cost2 - cost) / 1e-4, dw1[0, 0])


get_training_set()
check_gradient()
train()
predict()
