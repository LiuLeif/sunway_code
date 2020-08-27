#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-24 15:49
import tensorflow as tf
import numpy as np
import os

from tensorflow import keras
from tensorflow.keras import layers, losses, metrics, optimizers, models

from config import *
from models import *

x_train = np.load("./temp/train_x.npy")
y_train = np.load("./temp/train_y.npy")
x_train = np.squeeze(x_train)

tf.keras.backend.clear_session()

MODEL_PATH = "./temp/my"
if os.path.exists(MODEL_PATH):
    model = keras.models.load_model(MODEL_PATH)
else:
    model = cnn()

model.compile(
    optimizer="adam", loss="sparse_categorical_crossentropy", metrics=["accuracy"]
)
model.fit(x_train, y_train, batch_size=100, epochs=5, shuffle=True, verbose=1)

x_test = np.load("./temp/test_x.npy")
y_test = np.load("./temp/test_y.npy")
x_test = np.squeeze(x_test)

model.save(MODEL_PATH)

converter = tf.lite.TFLiteConverter.from_saved_model("./temp/my")
converter.optimizations = [tf.lite.Optimize.OPTIMIZE_FOR_SIZE]

tflite = converter.convert()
with open("./temp/output.tflite", "wb") as f:
    f.write(tflite)
    print("size of tflite:", len(tflite))

print("evalute for all")
model.evaluate(x_test, y_test, verbose=2)
for w in WORDS_TO_CHECK:
    x = np.load("./temp/test_" + w + "_x.npy")
    y = np.load("./temp/test_" + w + "_y.npy")
    x = np.squeeze(x)
    print("evalute for " + w + "")
    model.evaluate(x, y, verbose=2)
