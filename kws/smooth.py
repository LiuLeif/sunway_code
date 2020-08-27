#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-21 16:25

import collections
from config import *

N = SMOOTH_INTERVAL // INFERENCE_INTERVAL


class Smooth(object):
    def __init__(self):
        self.bucket = []

    def __call__(self, value):
        if N <= 1:
            return value

        self.bucket = self.bucket[1 - N :]
        self.bucket.append(value)

        if all(i == self.bucket[0] for i in self.bucket):
            return self.bucket[0]
        else:
            return "silent"


if __name__ == "__main__":
    N = 2
    smooth = Smooth()
    print(smooth(1))
    print(smooth(2))
    print(smooth(1))
    print(smooth(2))
    print(smooth(3))
