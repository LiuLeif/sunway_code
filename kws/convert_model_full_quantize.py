#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-19 09:00

import sys

if len(sys.argv) == 1:
    print("usage: convert_model.py <frozen_graph_file>")
    exit(1)

import tensorflow as tf
import numpy as np

converter = tf.compat.v1.lite.TFLiteConverter.from_frozen_graph(
    sys.argv[1], ["Mfcc"], ["labels_softmax"]
)

input_data = np.load("output2.npy")

def representative_dataset_gen():
    for i in range(10):
        yield [input_data]

converter.representative_dataset = representative_dataset_gen

converter.optimizations = [tf.lite.Optimize.DEFAULT]
converter.target_spec.supported_ops = [tf.lite.OpsSet.TFLITE_BUILTINS_INT8]

tflite_model = converter.convert()
with open("./output.tflite", "wb") as f:
    f.write(tflite_model)
    print("converted model saved to output.tflite")
    print("size of output.tflite:", len(tflite_model))
