#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-20 13:24
import sys

if len(sys.argv) == 1 or sys.argv[1] not in set(["train", "validation", "test"]):
    print("usage: convert_mfcc.py [train|validation|test]")
    exit(1)

import os
from mfcc import mfcc
import numpy as np
from config import *

mode = sys.argv[1]
DATA_DIR = "/home/sunway/download/speech_commands/" + mode + "/"

mapping = dict(zip(WORDS, range(len(WORDS))))
dirs = os.listdir(DATA_DIR)

X = []
Y = []

print(WORDS_TO_CHECK)

for d in dirs:
    category = 1  # unknown
    if d in mapping and d in WORDS_TO_CHECK:
        category = mapping[d]

    for f in os.listdir(DATA_DIR + d):
        X.append(mfcc(DATA_DIR + d + "/" + f))
        Y.append(category)

np.save("./temp/" + mode + "_x.npy", np.stack(X))
np.save("./temp/" + mode + "_y.npy", np.array(Y))

print("saved to ./temp/" + mode + "_x.npy and ./temp/" + mode + "_y.npy")
