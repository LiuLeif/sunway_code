#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-20 13:24
import os
from mfcc import mfcc
import numpy as np

DATA_DIR = "/home/sunway/download/speech_commands/test/"

words = [
    "silent",
    "unknown",
    "yes",
    "no",
    "up",
    "down",
    "left",
    "right",
    "on",
    "off",
    "stop",
    "go",
]

mapping = dict(zip(words, range(len(words))))
dirs = os.listdir(DATA_DIR)

X = []
Y = []

for d in dirs:
    category = 1  # unknown
    if d in mapping:
        category = mapping[d]
    for f in os.listdir(DATA_DIR + d):
        X.append(mfcc(DATA_DIR + d + "/" + f))
        Y.append(category)

np.save("/tmp/test_x.npy", np.stack(X))
np.save("/tmp/test_y.npy", np.array(Y))

print("saved to /tmp/test_x.npy and /tmp/test_y.npy")
