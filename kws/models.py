#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-26 09:52
import tensorflow as tf
import numpy as np
from config import *

from tensorflow import keras
from tensorflow.keras import layers, losses, metrics, optimizers, models


def cnn():
    inputs = keras.Input(shape=(FRAMES, DCT_COEFFICIENT_COUNT))
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
    outputs = layers.Dropout(0.1)(outputs)

    outputs = layers.Dense(128, activation="relu")(outputs)
    outputs = layers.Dense(len(WORDS), activation="softmax")(outputs)

    return keras.Model(inputs, outputs)


def dscnn():
    inputs = keras.Input(shape=(FRAMES, DCT_COEFFICIENT_COUNT))
    outputs = tf.expand_dims(inputs, axis=-1)

    outputs = layers.Conv2D(
        filters=172, kernel_size=[10, 4], strides=[2, 2], padding="same"
    )(outputs)
    outputs = layers.BatchNormalization()(outputs)
    outputs = layers.ReLU()(outputs)
    outputs = layers.Dropout(0.1)(outputs)

    # DepthwiseConv2D 1
    outputs = layers.DepthwiseConv2D(
        kernel_size=[3, 3], strides=[1, 1], padding="same"
    )(outputs)
    outputs = layers.Conv2D(
        filters=172, kernel_size=[1, 1], strides=[1, 1], padding="same"
    )(outputs)
    outputs = layers.BatchNormalization()(outputs)
    outputs = layers.ReLU()(outputs)
    outputs = layers.Dropout(0.2)(outputs)

    # DepthwiseConv2D 2
    outputs = layers.DepthwiseConv2D(
        kernel_size=[3, 3], strides=[1, 1], padding="same"
    )(outputs)
    outputs = layers.Conv2D(
        filters=172, kernel_size=[1, 1], strides=[1, 1], padding="same"
    )(outputs)
    outputs = layers.BatchNormalization()(outputs)
    outputs = layers.ReLU()(outputs)
    outputs = layers.Dropout(0.2)(outputs)

    # DepthwiseConv2D 3
    outputs = layers.DepthwiseConv2D(
        kernel_size=[3, 3], strides=[1, 1], padding="same"
    )(outputs)
    outputs = layers.Conv2D(
        filters=172, kernel_size=[1, 1], strides=[1, 1], padding="same"
    )(outputs)
    outputs = layers.BatchNormalization()(outputs)
    outputs = layers.ReLU()(outputs)
    outputs = layers.Dropout(0.2)(outputs)

    # DepthwiseConv2D 4
    outputs = layers.DepthwiseConv2D(
        kernel_size=[3, 3], strides=[1, 1], padding="same"
    )(outputs)
    outputs = layers.Conv2D(
        filters=172, kernel_size=[1, 1], strides=[1, 1], padding="same"
    )(outputs)
    outputs = layers.BatchNormalization()(outputs)
    outputs = layers.ReLU()(outputs)
    outputs = layers.Dropout(0.2)(outputs)

    outputs = layers.AveragePooling2D(
        pool_size=outputs.shape[1:3], strides=[2, 2], padding="valid"
    )(outputs)

    outputs = layers.Flatten()(outputs)
    outputs = layers.Dense(len(WORDS), activation="softmax")(outputs)

    return keras.Model(inputs, outputs)
