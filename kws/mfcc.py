#!/usr/bin/env python3

import numpy as np
import tensorflow as tf

from tensorflow.contrib.framework.python.ops import audio_ops as audio
from tensorflow.python.ops import io_ops

SAMPLE_RATE = 16000
CLIP_DURATION = 1000  # ms
WINDOW_SIZE_MS = 40  # ms
WINDOW_STRIDE_MS = 20  # ms
DCT_COEFFICIENT_COUNT = 10


class MFCC(object):
    """Handles loading, partitioning, and preparing audio training data."""

    def __init__(self,):
        self.prepare_processing_graph()

    def prepare_processing_graph(self):
        self.wav_filename_placeholder_ = tf.placeholder(tf.string, [])
        wav_loader = io_ops.read_file(self.wav_filename_placeholder_)
        wav_decoder = audio.decode_wav(
            wav_loader,
            desired_channels=1,
            desired_samples=SAMPLE_RATE * CLIP_DURATION / 1000,
        )

        window_size = int(SAMPLE_RATE * WINDOW_SIZE_MS / 1000)
        window_stride = int(SAMPLE_RATE * WINDOW_STRIDE_MS / 1000)

        spectrogram = audio.audio_spectrogram(
            wav_decoder.audio,
            window_size=window_size,
            stride=window_stride,
            magnitude_squared=True,
        )
        self.mfcc_ = audio.mfcc(
            spectrogram,
            wav_decoder.sample_rate,
            dct_coefficient_count=DCT_COEFFICIENT_COUNT,
        )

    def get(self, file_name):
        input_dict = {
            self.wav_filename_placeholder_: file_name,
        }
        with tf.Session() as sess:
            return sess.run(
                self.mfcc_, feed_dict={self.wav_filename_placeholder_: file_name,}
            )


if __name__ == "__main__":
    import sys
    if len(sys.argv) == 1:
        print("usage: mfcc <wav>")
        exit(1)
    mfcc = MFCC()
    x = mfcc.get(sys.argv[1])
    np.save("output", x)
    print("output to output.npy")
    print(x.shape)
