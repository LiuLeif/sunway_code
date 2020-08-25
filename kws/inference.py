#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-07-24 15:12
import numpy as np
import tensorflow as tf
import time
import os
import psutil
import sys
from config import *
from mfcc import mfcc


class Inference(object):
    def __init__(self):
        self.interpreter = tf.lite.Interpreter(model_path="./temp/output.tflite")
        self.interpreter.allocate_tensors()
        self.input_details = self.interpreter.get_input_details()
        self.output_details = self.interpreter.get_output_details()

    def __call__(self, input_data):
        self.interpreter.set_tensor(self.input_details[0]["index"], input_data)
        self.interpreter.invoke()
        output_data = self.interpreter.get_tensor(
            self.output_details[0]["index"]
        ).flatten()
        index = np.argmax(output_data)
        if output_data[index] >= MIN_PROBABILITY:
            return WORDS[index], output_data[index]
        else:
            return WORDS[0], 0


if __name__ == "__main__":
    if len(sys.argv) == 1:
        print("usage: inference <wav>")
        exit(1)
    inference = Inference()
    print(inference(mfcc(sys.argv[1])))
