#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-29 20:28

import pyaudio
from config import *
import numpy as np
import time
from progress import Progress


class MicStream(object):
    def __init__(self):
        self.buffer = np.array([], dtype=np.float32)
        p = pyaudio.PyAudio()
        self.stream = p.open(
            format=pyaudio.paInt16,
            channels=1,
            rate=SAMPLE_RATE,
            frames_per_buffer=CHUNK,
            input=True,
        )

    def __iter__(self):
        while True:
            data = self.stream.read(CHUNK)
            self.buffer = self.buffer[CHUNK - SAMPLES :]
            tmp = np.frombuffer(data, dtype=np.int16)
            self.buffer = np.append(self.buffer, tmp / 32767)
            if len(self.buffer) != SAMPLES:
                continue
            yield np.expand_dims(self.buffer, 1)


import tensorflow as tf
import numpy as np
from tensorflow.python.ops import gen_audio_ops as audio_ops
from tensorflow.python.ops import io_ops


class WavStream(object):
    def __init__(self, file_name):
        wav_loader = io_ops.read_file(file_name)
        wav_decoder = tf.audio.decode_wav(wav_loader, desired_channels=1)
        self.data = wav_decoder.audio
        self.length = self.data.shape[0]
        self.progress = Progress(self.length)

    def __iter__(self):
        for i in range(0, self.length - SAMPLES, CHUNK):
            self.progress(CHUNK)
            yield (self.data[i : i + SAMPLES])
            time.sleep(INFERENCE_INTERVAL / 1000)
