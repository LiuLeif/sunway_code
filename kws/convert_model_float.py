#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-19 09:00

import sys

if len(sys.argv) == 1:
    print("usage: convert_model.py <frozen_graph_file>")
    exit(1)

import tensorflow as tf

converter = tf.compat.v1.lite.TFLiteConverter.from_frozen_graph(
    sys.argv[1], ["Mfcc"], ["labels_softmax"]
)

tflite_model = converter.convert()
with open("/tmp/output.tflite", "wb") as f:
    f.write(tflite_model)
    print("converted model saved to /tmp/output.tflite")
    print("size of output.tflite:", len(tflite_model))
