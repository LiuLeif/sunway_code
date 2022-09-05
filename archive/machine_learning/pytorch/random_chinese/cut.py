#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2021-08-25 11:18
import jieba

s = open("bai_lu_yuan.txt", "r").readlines()
for l in s:
    seg_list = jieba.cut(l, cut_all=False)
    print(" ".join(seg_list))

