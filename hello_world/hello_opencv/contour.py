#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-12-28 14:25
import cv2
import numpy as np

# findContours 要求输入图像是 grayscale
test_img = np.ones((400, 400), np.uint8) * 255
# drawContours 相当于按照给定的顶点_按顺序_画一个多边形
cv2.drawContours(
    test_img,
    np.array([[[10, 10], [10, 50], [50, 50], [50, 10]]]),
    -1,
    (0),
    3,
)

# findContours 会根据 edge detection 的结果返回多个 contour
contours, hierarchy = cv2.findContours(test_img, cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE)
for c in contours:
    print(c)

cv2.drawContours(
    test_img,
    contours,
    -1,
    (200),
    3,
)
cv2.imshow("", test_img)
while cv2.waitKey(1) != ord("q"):
    pass
