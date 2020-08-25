#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-24 15:49
import tensorflow as tf
import numpy as np
import os

from tensorflow import keras
from tensorflow.keras.datasets import mnist
from tensorflow.keras import layers, losses, metrics, optimizers, models

from config import *

# tf.config.set_visible_devices([], "GPU")

x_train = np.load("./temp/train_x.npy")
y_train = np.load("./temp/train_y.npy")
x_train = np.squeeze(x_train)

tf.keras.backend.clear_session()

MODEL_PATH = "./temp/my"
if os.path.exists(MODEL_PATH):
    model = keras.models.load_model(MODEL_PATH)
else:
    inputs = keras.Input(shape=(49, 10))
    outputs = tf.expand_dims(inputs, axis=-1)

    outputs = layers.Conv2D(filters=64, kernel_size=[10, 4], strides=[1, 1])(outputs)
    outputs = layers.BatchNormalization()(outputs)
    outputs = layers.ReLU()(outputs)
    outputs = layers.Dropout(0.5)(outputs)

    outputs = layers.Conv2D(filters=48, kernel_size=[10, 4], strides=[2, 1])(outputs)
    outputs = layers.BatchNormalization()(outputs)
    outputs = layers.ReLU()(outputs)
    outputs = layers.Dropout(0.5)(outputs)

    outputs = layers.Flatten()(outputs)

    outputs = layers.Dense(16)(outputs)
    outputs = layers.BatchNormalization()(outputs)
    outputs = layers.ReLU()(outputs)
    outputs = layers.Dropout(0.5)(outputs)

    outputs = layers.Dense(128, activation="relu")(outputs)
    outputs = layers.Dense(12, activation="softmax")(outputs)

    model = keras.Model(inputs, outputs)

model.compile(
    optimizer="adam", loss="sparse_categorical_crossentropy", metrics=["accuracy"]
)
model.fit(x_train, y_train, batch_size=200, epochs=20, shuffle=True, verbose=1)

x_test = np.load("./temp/test_x.npy")
y_test = np.load("./temp/test_y.npy")
x_test = np.squeeze(x_test)

print("evalute for all")
model.evaluate(x_test, y_test, verbose=2)
for w in WORDS_TO_CHECK:
    x = np.load("./temp/test_" + w + "_x.npy")
    y = np.load("./temp/test_" + w + "_y.npy")
    x = np.squeeze(x)
    print("evalute for " + w + "")
    model.evaluate(x, y, verbose=2)

model.save(MODEL_PATH)
