#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-07-24 15:12
import numpy as np
import tensorflow as tf
import time
import os
import psutil
from config import *

interpreter = tf.lite.Interpreter(model_path="./model/output.tflite")
interpreter.allocate_tensors()

input_details = interpreter.get_input_details()
output_details = interpreter.get_output_details()

for w in sorted(WORDS):
    X = np.load("./temp/test_" + w + "_x.npy")
    Y = np.load("./temp/test_" + w + "_y.npy")

    print("evalute for " + w)
    prediction = []
    for x in X:
        interpreter.set_tensor(input_details[0]["index"], x)
        interpreter.invoke()
        output_data = interpreter.get_tensor(output_details[0]["index"]).flatten()
        y = np.argmax(output_data)
        prediction.append(y)

    print((np.array(prediction) == Y).sum() / len(Y))
