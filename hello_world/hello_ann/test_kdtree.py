#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-12-23 18:29
import numpy as np
from sklearn.neighbors import KDTree
import time

np.random.seed(0)
X = np.random.randn(1000, 1000)

tree = KDTree(X, leaf_size=1000)
before = time.time()
dist, index = tree.query(X[:1], k=3)
print(index)  # 3个最近的索引
print(int((time.time() - before) * 1000))
