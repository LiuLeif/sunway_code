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
from sklearn.linear_model.logistic import LogisticRegression

POLY_FEATURES = 4


def draw_decision_boundary(classifier):
    xx, yy = np.meshgrid(
        np.arange(-1.5, 2.5, 0.02), np.arange(-1.5, 2.5, 0.02))
    X = np.c_[xx.ravel(), yy.ravel()]

    poly = preprocessing.PolynomialFeatures(POLY_FEATURES, include_bias=False)
    X = poly.fit_transform(X)

    Y = classifier.predict(X)
    Y = Y.reshape(xx.shape)
    cm = ListedColormap(['#FF0000', '#0000FF'])
    plt.contour(xx, yy, Y, cmap=cm)


def get_training_set():
    X, Y = make_moons(n_samples=1000, noise=0.2)
    # X, Y = make_circles(n_samples=1000, noise=0.2, factor=0.5)
    Y = Y.reshape(1000, 1)
    poly = preprocessing.PolynomialFeatures(POLY_FEATURES, include_bias=False)
    X = poly.fit_transform(X, Y)
    return X, Y


def predict():
    X, Y = get_training_set()
    classifier = LogisticRegression()
    classifier.fit(X, Y)

    cm = ListedColormap(['#FF0000', '#0000FF'])
    plt.scatter(x=X[:, 0], y=X[:, 1], c=Y[:, 0], cmap=cm)
    draw_decision_boundary(classifier)
    plt.show()


predict()
