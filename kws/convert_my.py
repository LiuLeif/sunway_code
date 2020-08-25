#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-24 16:20
import tensorflow as tf

converter = tf.lite.TFLiteConverter.from_saved_model("./temp/my")
converter.optimizations = [tf.lite.Optimize.OPTIMIZE_FOR_SIZE]

tflite = converter.convert()
with open("./temp/output.tflite", "wb") as f:
    f.write(tflite)
    print("size of tflite:", len(tflite))
