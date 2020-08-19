#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-07-24 15:12
import numpy as np
import tensorflow as tf

interpreter = tf.lite.Interpreter(model_path="output.tflite")
interpreter.allocate_tensors()

input_details = interpreter.get_input_details()
output_details = interpreter.get_output_details()

input_data = np.load("output.npy")
interpreter.set_tensor(input_details[0]["index"], input_data)
interpreter.invoke()

output_data = interpreter.get_tensor(output_details[0]["index"]).flatten()
# '[silence],[unknown],yes,no,up,down,left,right,on,off,stop,go'
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
