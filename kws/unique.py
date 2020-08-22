#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-22 19:13
from config import *


class Unique(object):
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
    unique = Unique()
    print(unique("right"))
    print(unique("left"))
    print(unique("unknown"))
    print(unique("silent"))
    print(unique("left"))
    print(unique("right"))
