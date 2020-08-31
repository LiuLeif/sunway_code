#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-28 11:58
import tensorflow as tf
import numpy as np
from config import *


def up_sample(data):
    data = tf.concat((data, data), axis=1)
    data = tf.reshape(data, (-1, 1))
    offset = np.random.randint(0, SMOOTH_SAMPLES * 2)
    return data[offset : offset + SAMPLES]


if __name__ == "__main__":
    import sys

    from tensorflow.python.ops import gen_audio_ops as audio_ops
    from tensorflow.python.ops import io_ops

    if len(sys.argv) == 1:
        print("usage: up_sample <wav>")
        exit(1)

    wav_loader = io_ops.read_file(sys.argv[1])
    wav_decoder = tf.audio.decode_wav(
        wav_loader, desired_channels=1, desired_samples=SAMPLES,
    )
    data = wav_decoder.audio
    data = up_sample(data)

    wav_file = tf.audio.encode_wav(data, SAMPLE_RATE)
    with open("output.wav", "wb") as f:
        f.write(wav_file.numpy())
