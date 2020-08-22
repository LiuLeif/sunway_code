#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-07-24 15:12
import numpy as np
import tensorflow as tf
import time
import os
import psutil
from config import *

WORDS = [
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


class Inference(object):
    def __init__(self):
        self.interpreter = tf.lite.Interpreter(model_path="./temp/output.tflite")
        self.interpreter.allocate_tensors()
        self.input_details = self.interpreter.get_input_details()
        self.output_details = self.interpreter.get_output_details()

    def run(self, input_data):
        self.interpreter.set_tensor(self.input_details[0]["index"], input_data)
        self.interpreter.invoke()
        output_data = self.interpreter.get_tensor(
            self.output_details[0]["index"]
        ).flatten()
        index = np.argmax(output_data)
        if output_data[index] > MIN_PROBABILITY:
            return WORDS[index]
        else:
            return WORDS[0]


if __name__ == "__main__":
    inference = Inference()
    print(inference.run(np.load("./temp/output.npy")))
