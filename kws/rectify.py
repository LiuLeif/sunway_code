#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-22 19:13
from config import *


class FastResponseRectify(object):
    def __init__(self):
        self.prev = None
        self.count = 0
        self.response_count = 800 // INFERENCE_INTERVAL

    def __call__(self, value):
        self.count += 1

        if value == "silent":
            value = "unknown"

        if self.prev == None or self.prev != value:
            self.prev = value
            if value != "unknown" and self.count >= self.response_count:
                self.count = 0
                return value
        else:
            return None


class SilenceRectify(object):
    def __init__(self):
        self.prev = None

    # 0 0 1 1 0 1 2 1 0    smoothed sequence
    #     1         1      silent
    #   0         2        output
    def __call__(self, value):
        prev, self.prev = self.prev, value
        if value != "silent":
            return None
        if prev == "unknown" or prev == "silent":
            return None
        return prev


if __name__ == "__main__":
    rectify = Rectify()
    print(rectify("right"))
    print(rectify("left"))
    print(rectify("unknown"))
    print(rectify("silent"))
    print(rectify("left"))
    print(rectify("right"))
