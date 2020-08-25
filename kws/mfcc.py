#!/usr/bin/env python3

import numpy as np
import tensorflow as tf
from config import *

from tensorflow.python.ops import gen_audio_ops as audio_ops
from tensorflow.python.ops import io_ops


def mfcc_data(data):
    window_size = int(SAMPLE_RATE * WINDOW_SIZE_MS / 1000)
    window_stride = int(SAMPLE_RATE * WINDOW_STRIDE_MS / 1000)

    spectrogram = audio_ops.audio_spectrogram(
        data, window_size=window_size, stride=window_stride, magnitude_squared=True,
    )
    return audio_ops.mfcc(
        spectrogram, SAMPLE_RATE, dct_coefficient_count=DCT_COEFFICIENT_COUNT,
    )


def apply_background_noise(data, noise):
    data += noise
    return data


# max_shift: ms
def apply_time_shift(data, max_shift):
    shift_length = np.random.randint(-max_shift, max_shift)
    if shift_length > 0:
        padding = np.array([[shift_length, 0], [0, 0]])
        offset = [0, 0]
    else:
        padding = np.array([[0, -shift_length], [0, 0]])
        offset = [-shift_length, 0]
    data = tf.pad(data, padding)
    data = tf.slice(data, offset, [SAMPLES, -1])
    return data


def mfcc(file_name, time_shift=0, noise=None):
    wav_loader = io_ops.read_file(file_name)
    wav_decoder = tf.audio.decode_wav(
        wav_loader, desired_channels=1, desired_samples=SAMPLES,
    )
    data = wav_decoder.audio
    if time_shift != 0:
        data = apply_time_shift(data, time_shift)
    if noise is not None:
        data = apply_background_noise(data, noise)
    return mfcc_data(data)


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
