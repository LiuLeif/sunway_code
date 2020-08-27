#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-20 13:24
import sys
import collections

if len(sys.argv) == 1 or sys.argv[1] not in set(["train", "validation", "test"]):
    print("usage: convert_mfcc.py [train|validation|test]")
    exit(1)

import os
from mfcc import *
from noise import Noise
import numpy as np

from config import *

mode = sys.argv[1]
DATA_DIR = "/home/sunway/download/speech_commands/" + mode + "/"

mapping = dict(zip(WORDS, range(len(WORDS))))
dirs = os.listdir(DATA_DIR)

X = []
Y = []

SEPERATE_X = collections.defaultdict(list)
SEPERATE_Y = collections.defaultdict(list)

print(WORDS_TO_CHECK)

noise = Noise()

for d in dirs:
    if d not in WORDS:
        category = 1
    else:
        if d not in WORDS_TO_CHECK:
            continue
        else:
            category = mapping[d]

    print(d, WORDS[category])
    seperate_x = SEPERATE_X[WORDS[category]]
    seperate_y = SEPERATE_Y[WORDS[category]]

    def do_mfcc(f, time_shift=0, noise=None):
        fingerprint = mfcc(f, time_shift, noise)
        X.append(fingerprint)
        Y.append(category)
        if mode != "train":
            seperate_x.append(fingerprint)
            seperate_y.append(category)

    for f in os.listdir(DATA_DIR + d):
        if mode == "train":
            for _ in range(10):
                do_mfcc(DATA_DIR + d + "/" + f, time_shift=200, noise=noise())
        else:
            do_mfcc(DATA_DIR + d + "/" + f)

np.save("./temp/" + mode + "_x.npy", np.stack(X))
np.save("./temp/" + mode + "_y.npy", np.array(Y))

if mode != "train":
    for (k, v) in SEPERATE_X.items():
        np.save("./temp/" + mode + "_" + k + "_x.npy", np.stack(v))
        np.save("./temp/" + mode + "_" + k + "_y.npy", np.array(SEPERATE_Y[k]))

print("saved to ./temp/" + mode)
