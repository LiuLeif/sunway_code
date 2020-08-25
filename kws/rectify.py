#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-22 19:13
from config import *


class Rectify(object):
    def __init__(self):
        self.count = 0
        self.pending = None
        self.bucket = 0

    def __call__(self, value):
        self.count += 1
        if value == "unknown" or value == "silent":
            self.bucket |= 1
            return None

        self.pending = value

        if self.count >= RESPONSE_INTERVAL // INFERENCE_INTERVAL:
            self.bucket |= 2

        if self.bucket == 3:
            self.count = 0
            self.bucket = 0
            return self.pending

        return None


if __name__ == "__main__":
    rectify = Rectify()
    print(rectify("right"))
    print(rectify("left"))
    print(rectify("unknown"))
    print(rectify("silent"))
    print(rectify("left"))
    print(rectify("right"))
