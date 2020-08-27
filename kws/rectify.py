#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-22 19:13
from config import *


class Rectify(object):
    def __init__(self):
        self.count = 0
        self.pending = None
        self.flag = 0

    def __call__(self, value):
        self.count += 1
        if value == "silent" or value == "unknown":
            self.flag |= 1
            return None

        self.pending = value

        if self.count >= RESPONSE_INTERVAL // INFERENCE_INTERVAL:
            self.flag |= 2

        if self.flag == 3:
            self.count = 0
            self.flag = 0
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
