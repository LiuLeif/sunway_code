#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2021-10-14 15:43
import caffe
import numpy as np

net = caffe.Net("test.prototxt", "test_solver_iter_1000.caffemodel", caffe.TEST)
import ipdb; ipdb.set_trace()
net.blobs["data"].data[...] = np.array([10])
print(net.forward())
