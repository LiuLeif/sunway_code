#!/usr/bin/env python3
import numpy as np
import cv2
import glob

# prepare object points, like (0,0,0), (1,0,0), (2,0,0) ....,(6,5,0)
# objp = np.zeros((6 * 9, 3), np.float32)
# objp[:, :2] = np.mgrid[0:9, 0:6].T.reshape(-1, 2)

objp = np.concatenate(
    (
        np.array(np.meshgrid(np.arange(9), np.arange(6)))
        .transpose(1, 2, 0)
        .reshape(-1, 2),
        np.zeros((54, 1)),
    ),
    axis=1,
).astype(np.float32)

obj_points = []  # 3d point in real world space
img_points = []  # 2d points in image plane.

images = glob.glob("images/*.jpg")  # read a series of images

for image in images:
    img = cv2.imread(image)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)  # Convert the image to gray

    # Find the chess board corners
    ret, corners = cv2.findChessboardCorners(gray, (9, 6), None)

    if ret:
        obj_points.append(objp)
        img_points.append(corners)

_, M, dist, rvecs, tvecs = cv2.calibrateCamera(
    obj_points, img_points, gray.shape[::-1], None, None
)

print("intrinsic matrix=\n", M)
# print("distortion coefficients=\n", dist)
# print("rotation vector for each image=", *rvecs, sep="\n")
# print("translation vector for each image=", *tvecs, sep="\n")
