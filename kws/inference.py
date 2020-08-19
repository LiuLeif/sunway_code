#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-07-24 15:12
import numpy as np
import tensorflow as tf
import time
import os
import psutil

process = psutil.Process(os.getpid())

interpreter = tf.lite.Interpreter(model_path="output.tflite")
print(">allocate tensors")
time.sleep(2);
before = process.memory_info().vms
print(process.memory_info())

print(">>>allocate tensors")
interpreter.allocate_tensors()
time.sleep(2);

after = process.memory_info().vms
print(process.memory_info())
print("<<<allocate tensors")

input_details = interpreter.get_input_details()
output_details = interpreter.get_output_details()

input_data = np.load("output.npy")

interpreter.set_tensor(input_details[0]["index"], input_data)
interpreter.invoke()

last = process.memory_info().vms
print(process.memory_info())

print("result:", (after - before) / 1024, (last - after) / 1024)
output_data = interpreter.get_tensor(output_details[0]["index"]).flatten()

words = [
    "silent",
    "unknown",
    "yes",
    "no",
    "up",
    "down",
    "left",
    "right",
    "on",
    "off",
    "stop",
    "go",
]
print(output_data)
print(words[np.argmax(output_data)])
