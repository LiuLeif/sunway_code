#!/usr/bin/env python3
import numpy as np
import cv2
import glob

# prepare object points, like (0,0,0), (1,0,0), (2,0,0) ....,(6,5,0)
# objp = np.zeros((6 * 9, 3), np.float32)
# objp[:, :2] = np.mgrid[0:9, 0:6].T.reshape(-1, 2)
CHECKERBOARD = (6, 8)
criteria = (cv2.TERM_CRITERIA_EPS + cv2.TERM_CRITERIA_MAX_ITER, 30, 0.001)
objp = np.concatenate(
    (
        np.array(np.meshgrid(np.arange(8), np.arange(6)))
        .transpose(1, 2, 0)
        .reshape(-1, 2),
        np.zeros((48, 1)),
    ),
    axis=1,
).astype(np.float32)

obj_points = []  # 3d point in real world space
img_points = []  # 2d points in image plane.

images = glob.glob("images/*.png")  # read a series of images

for image in images:
    img = cv2.imread(image)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)  # Convert the image to gray

    # Find the chess board corners
    ret, corners = cv2.findChessboardCorners(gray, (8, 6), None)
    print(ret)

    if ret:
        obj_points.append(objp)
        cv2.cornerSubPix(gray, corners, (11, 11), (-1, -1), criteria)
        img_points.append(corners)
        img = cv2.drawChessboardCorners(img, CHECKERBOARD, corners, ret)
        cv2.imshow("img", img)
        cv2.waitKey(300)

_, M, dist, rvecs, tvecs = cv2.calibrateCamera(
    obj_points, img_points, gray.shape[::-1], None, None
)

print("intrinsic matrix=\n", M)
# print("distortion coefficients=\n", dist)
# print("rotation vector for each image=", *rvecs, sep="\n")
# print("translation vector for each image=", *tvecs, sep="\n")
