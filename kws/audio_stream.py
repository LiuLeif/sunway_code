#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-21 11:57
import sys
from stream import *

if len(sys.argv) != 2:
    print("usage: audio_stream  <mic|xxx.wav>")
    exit(1)

if sys.argv[1] == "mic":
    stream = MicStream()
else:
    stream = WavStream(sys.argv[1])

import numpy as np
import tensorflow as tf
from inference import Inference
from smooth import *
from rectify import *
from config import *
from log import *
from mfcc import MFCC

inference = Inference()
smooth = MaximumSmooth()
rectify = SilenceRectify()

for (i, data) in enumerate(stream):
    x = MFCC.decode_pcm(data)
    output, prob = inference(x)
    get_logger().info("inf: %6d %s %f" % (i, output, prob))
    output = smooth(output)
    get_logger().info("    smo: %s" % (output))
    output = rectify(output)
    if output:
        get_logger().info("        result: %-6s:%f" % (output, prob))
        wav_file = tf.audio.encode_wav(data, SAMPLE_RATE)
        with open("./logs/%s_%d.wav" % (output, i), "wb") as f:
            f.write(wav_file.numpy())
        print("%-6s %-6f %d" % (output, prob, i))
