#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-30 17:57
import sys


class Progress(object):
    def __init__(self, total):
        self.total = total
        self.i = 0

    def __call__(self, step):
        self.i += step
        progress = self.i * 100 // self.total
        sys.stdout.write(("" * (100 - progress)) + ("\r [ %d" % progress + "% ] "))
