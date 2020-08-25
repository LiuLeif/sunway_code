#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-25 17:15
import os
import numpy as np
import tensorflow as tf
from tensorflow.python.ops import io_ops
from config import *

NOISE_DIR = "/home/sunway/download/speech_commands/_background_noise_/"


class Noise(object):
    def __init__(self):
        self.data = []
        for f in os.listdir(NOISE_DIR):
            wav_loader = io_ops.read_file(NOISE_DIR + f)
            wav_decoder = tf.audio.decode_wav(wav_loader, desired_channels=1)
            self.data.append(wav_decoder.audio * 0.1)

    def __call__(self):
        index = np.random.randint(0, len(self.data))
        offset = np.random.randint(0, self.data[index].shape[0] - SAMPLES)
        return self.data[index][offset : offset + SAMPLES]


if __name__ == "__main__":
    noise = Noise()
    print(noise())
    print(noise())
