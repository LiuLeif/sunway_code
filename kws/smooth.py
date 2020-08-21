#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-21 16:25

import collections

N = 50  # 1s/20ms = 50


class Smooth(object):
    def __init__(self):
        self.bucket = []
        self.counter = collections.defaultdict(int)

    def feed(self, value):
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
    N = 3
    smooth = Smooth()
    print(smooth.feed(1))
    print(smooth.feed(2))
    print(smooth.feed(1))
    print(smooth.feed(2))
    print(smooth.feed(3))
