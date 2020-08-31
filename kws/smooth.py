#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-21 16:25

import collections
from config import *

N = SMOOTH_INTERVAL // INFERENCE_INTERVAL


class ContiousSmooth(object):
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
            return "unknown"


class MaximumSmooth(object):
    def __init__(self):
        self.bucket = []
        self.counter = collections.defaultdict(int)

    def __call__(self, value):
        if N <= 1:
            return value

        if len(self.bucket) == N:
            self.counter[self.bucket[0]] -= 1

        self.bucket = self.bucket[1 - N :]
        self.bucket.append(value)
        self.counter[value] += 1

        max_key = 0
        max_val = 0

        for (k, v) in self.counter.items():
            if v >= max_val:
                max_val = v
                max_key = k
        return max_key


if __name__ == "__main__":
    N = 2
    smooth = ContiousSmooth()
    print(smooth(1))
    print(smooth(2))
    print(smooth(1))
    print(smooth(2))
    print(smooth(3))
