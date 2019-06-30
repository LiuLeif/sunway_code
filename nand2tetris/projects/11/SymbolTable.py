#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2018-09-13 16:44

from collections import defaultdict


class SymbolTable:
    def __init__(self):
        self.symbols = {}
        self.indexs = defaultdict(int)

    def reset(self):
        self.symbols = {}
        self.indexs = defaultdict(int)
        
    def mark_symbol(self, return_t, name, type_):
        self.symbols[name] = (name, return_t, type_, self.indexs[type_])
        self.indexs[type_] += 1
        print(name, type_, return_t)

    def contains(self, name):
        return name in self.symbols

    def get_type(self, name):
        return self.symbols[name][2]

    def get_index(self, name):
        return self.symbols[name][3]
    
    def get_symbol_type(self, name):
        return self.symbols[name][1]

    def get_count(self, category):
        return self.indexs[category]
