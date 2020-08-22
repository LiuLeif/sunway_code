#!/usr/bin/env python3

import numpy as np
import tensorflow as tf
from config import *

from tensorflow.python.ops import gen_audio_ops as audio_ops
from tensorflow.python.ops import io_ops

SAMPLE_RATE = 16000
WINDOW_SIZE_MS = 40  # ms
WINDOW_STRIDE_MS = 20  # ms
DCT_COEFFICIENT_COUNT = 10


def mfcc_data(data):
    window_size = int(SAMPLE_RATE * WINDOW_SIZE_MS / 1000)
    window_stride = int(SAMPLE_RATE * WINDOW_STRIDE_MS / 1000)

    spectrogram = audio_ops.audio_spectrogram(
        data, window_size=window_size, stride=window_stride, magnitude_squared=True,
    )
    return audio_ops.mfcc(
        spectrogram, SAMPLE_RATE, dct_coefficient_count=DCT_COEFFICIENT_COUNT,
    )


def mfcc(file_name):
    wav_loader = io_ops.read_file(file_name)
    wav_decoder = tf.audio.decode_wav(
        wav_loader,
        desired_channels=1,
        desired_samples=SAMPLE_RATE * CLIP_DURATION / 1000,
    )
    return mfcc_data(wav_decoder.audio)


def export_c(data):
    # float data[] = {}
    with open("./temp/output.cc", "w") as f:
        f.write("float data[] = {\n")
        for x in data:
            f.write(str(x) + ",")
        f.write("\n};")
    print("output to ./temp/output.cc")


def export_python(data):
    np.save("./temp/output.npy", data)
    print("output to ./temp/output.npy")


if __name__ == "__main__":
    import sys

    if len(sys.argv) == 1:
        print("usage: mfcc <wav>")
        exit(1)

    x = mfcc(sys.argv[1])
    print(x)
    print(x.shape)

    export_python(x)
    export_c(x.numpy().flatten())
