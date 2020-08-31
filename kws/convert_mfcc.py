#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-20 13:24
import sys
import collections

import os
from mfcc import *
from noise import Noise
import numpy as np

from config import *
from progress import Progress


class MFCCConvertor(object):
    def __init__(self, mode):
        self.mode = mode
        self.data_dir = "/home/sunway/download/speech_commands/" + mode + "/"
        self.mapping = dict(zip(WORDS, range(len(WORDS))))
        self.pending = []
        self.noise = Noise()
        if not os.path.exists("./temp"):
            os.mkdir("./temp")

    def scan(self):
        for d in os.listdir(self.data_dir):
            category = self.mapping.get(d, 0)
            for f in os.listdir(self.data_dir + d):
                self.pending.append((self.data_dir + d + "/" + f, category))

        counter = collections.Counter([t[1] for t in self.pending])
        print("datasets:", dict([(WORDS[k], v) for (k, v) in counter.items()]))

    def convert(self, required_size):
        print("required:", required_size)

        X = []
        Y = []

        SEPERATE_X = collections.defaultdict(list)
        SEPERATE_Y = collections.defaultdict(list)

        required_size = [required_size[WORDS[i]] for i in range(len(WORDS))]
        progress = Progress(sum(required_size))
        while any([x > 0 for x in required_size]):
            for f, category in self.pending:
                if required_size[category] > 0:
                    required_size[category] -= 1
                    fingerprint = MFCC.decode_wav(f, SMOOTH_INTERVAL, self.noise())
                    X.append(fingerprint)
                    Y.append(category)
                    SEPERATE_X[category].append(fingerprint)
                    SEPERATE_Y[category].append(category)
                    progress(1)

        np.save("./temp/" + self.mode + "_x.npy", np.stack(X))
        np.save("./temp/" + self.mode + "_y.npy", np.array(Y))

        for (k, v) in SEPERATE_X.items():
            np.save("./temp/" + mode + "_" + WORDS[k] + "_x.npy", np.stack(v))
            np.save(
                "./temp/" + mode + "_" + WORDS[k] + "_y.npy", np.array(SEPERATE_Y[k])
            )


if __name__ == "__main__":
    if len(sys.argv) == 1 or sys.argv[1] not in set(["train", "validation", "test"]):
        print("usage: convert_mfcc.py [train|validation|test]")
        exit(1)

    mode = sys.argv[1]
    convertor = MFCCConvertor(mode)
    convertor.scan()
    if mode == "train":
        convertor.convert(
            required_size={"unknown": 80000, "silent": 20000, "right": 20000}
        )
    else:
        convertor.convert(
            required_size={"unknown": 10000, "silent": 2500, "right": 2500}
        )
