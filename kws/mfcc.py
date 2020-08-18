# Copyright 2017 The TensorFlow Authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ==============================================================================
"""Model definitions for simple speech recognition.

"""

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
            ).flatten()


if __name__ == "__main__":
    mfcc = MFCC()
    x = mfcc.get("/tmp/test.wav")
    print(x.shape)
