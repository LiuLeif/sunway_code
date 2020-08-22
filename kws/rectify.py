#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-22 19:13
from config import *


class Rectify(object):
    def __init__(self):
        self.count = 0
        self.prev = "silent"

    def __call__(self, value):
        self.count += 1
        if value == "unknown":
            value = "silent"
        if self.prev == value:
            return None
        if self.prev != "silent":
            self.prev = value
            return None

        self.prev = value
        if self.count >= RESPONSE_INTERVAL // INFERENCE_INTERVAL:
            self.count = 0
            return value
        return None


if __name__ == "__main__":
    rectify = Rectify()
    print(rectify("right"))
    print(rectify("left"))
    print(rectify("unknown"))
    print(rectify("silent"))
    print(rectify("left"))
    print(rectify("right"))
