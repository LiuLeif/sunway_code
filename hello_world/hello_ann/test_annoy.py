#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-12-23 18:09

import time
from annoy import AnnoyIndex
import numpy as np

index = AnnoyIndex(5, "angular")
index.add_item(0, np.array([1, 2, 3, 4, 5]))
index.add_item(1, np.array([2, 3, 4, 4, 5]))
index.add_item(3, np.array([1, 2, 3, 4, 5]))
index.build(10)
x = index.get_nns_by_item(0, 2, include_distances=True)
# ([0, 3], [0.0, 0.0])
print(x)
x = index.get_nns_by_vector(np.array([2, 3, 4, 4, 5]), 1, include_distances=True)
# ([1], [0.0])
print(x)
# [2.0, 3.0, 4.0, 4.0, 5.0]
print(index.get_item_vector(1))
# [0.0, 0.0, 0.0, 0.0, 0.0]
print(index.get_item_vector(2))
# 4
print(index.get_n_items())
# 0.1838257759809494
print(index.get_distance(0, 1))

index.save("test.ann")
index = AnnoyIndex(5, "angular")
index.load("test.ann")
# 4
print(index.get_n_items())


index = AnnoyIndex(1000, "angular")
for i in range(10000):
    index.add_item(i, np.random.randn(1000))
index.build(10)
before = time.time()
x = index.get_nns_by_item(0, 4, include_distances=False)
# [0, 1438, 4199, 265]
print(x)
# 0
print(int((time.time() - before) * 1000))
