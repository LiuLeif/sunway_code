#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-05-02 11:09
import numpy as np
import matplotlib.pyplot as plt
import tensorflow as tf
from sklearn import preprocessing

LEARNING_RATE = 0.01


def MSE(A, B):
    return tf.reduce_mean(tf.square(tf.subtract(A, B)))


X_train, Y_train = 0, 0


def get_training_set():
    global X_train, Y_train
    data = np.loadtxt("data.txt", delimiter=",")
    X_train = data[:, 0].reshape(-1, 1)
    X_train = preprocessing.scale(X_train)
    Y_train = data[:, 1].reshape(-1, 1)
    Y_train = preprocessing.scale(Y_train)


get_training_set()

X = tf.placeholder(tf.float32, X_train.shape)
Y = tf.placeholder(tf.float32, Y_train.shape)

W = tf.get_variable(
    "weights", shape=(1, 1), initializer=tf.random_normal_initializer())
B = tf.get_variable(
    "bias", shape=(1, 1), initializer=tf.random_normal_initializer())

y_hat = tf.matmul(X, W) + B
cost = MSE(y_hat, Y)

optimizer = tf.train.GradientDescentOptimizer(LEARNING_RATE).minimize(cost)
with tf.Session() as sess:
    sess.run(tf.global_variables_initializer())
    for i in range(1000):
        _, loss = sess.run(
            [optimizer, cost], feed_dict={
                X: X_train,
                Y: Y_train,
            })
        print("#training #%d, %f" % (i, loss))
    w = sess.run(W)
    b = sess.run(B)
    plt.scatter(x=X_train[:, 0], y=Y_train[:, 0])
    plt.plot(X_train, np.matmul(X_train, w) + b, "r")
    plt.show()
