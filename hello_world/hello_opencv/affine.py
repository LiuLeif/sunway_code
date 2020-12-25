#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-12-24 22:09
import cv2
import numpy as np

img = cv2.imread("lena.jpg")
src = np.float32([[0, 0], [0, 1], [1, 0]])
dst = np.float32([[0, 10], [0, 11], [1, 10]])

#---------- getAffinetransform
M = cv2.getAffineTransform(src, dst)
print(M)
row, col = img.shape[:2]

# ---------- warpAffine
# cv2 的第一维为 col, 而 numpy 第一维是 row
img = cv2.warpAffine(img, M, img.shape[:2][::-1])
cv2.imshow("", img)

# ---------- transform
print(cv2.transform(np.expand_dims(src, 0), M))
print(cv2.transform(np.expand_dims(dst, 0), cv2.invertAffineTransform(M)))
while cv2.waitKey(1) != ord('q'):
    pass
