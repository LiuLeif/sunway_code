#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-07-24 15:12
import numpy as np
import tensorflow as tf
import time
import os
import psutil

process = psutil.Process(os.getpid())

interpreter = tf.lite.Interpreter(model_path="/tmp/output.tflite")
interpreter.allocate_tensors()

input_details = interpreter.get_input_details()
output_details = interpreter.get_output_details()

input_data = np.load("/tmp/test_x.npy")
true_data = np.load("/tmp/test_y.npy")

prediction = []
for x in input_data:
    interpreter.set_tensor(input_details[0]["index"], x)
    interpreter.invoke()
    output_data = interpreter.get_tensor(output_details[0]["index"]).flatten()
    y = np.argmax(output_data)
    prediction.append(y)

print((np.array(prediction) == true_data).sum() / len(true_data))
