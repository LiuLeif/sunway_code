#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2021-12-30 15:06
import cv2
import torch
import torchvision

import numpy as np
import torch.nn.functional as F

from scnn import SCNN

resize_shape = (800, 288)

mean = (0.485, 0.456, 0.406)
std = (0.229, 0.224, 0.225)


def resize(img, size):
    img = cv2.resize(img, size, interpolation=cv2.INTER_CUBIC)
    return img


def to_tensor(img):
    img = img.transpose(2, 0, 1)
    img = torch.from_numpy(img).type(torch.float) / 255.0
    return img


def normalize(img, mean, std):
    return torchvision.transforms.Normalize(mean, std)(img)


# download hello_tusimple.pth from
# https://drive.google.com/open?id=1IwEenTekMt-t6Yr5WJU9_kv4d_Pegd_Q
net = SCNN(input_size=resize_shape, pretrained=False)
save_dict = torch.load("hello_tusimple.pth")
net.load_state_dict(save_dict["net"])
net.eval()

with torch.no_grad():
    img = cv2.imread("hello_tusimple.jpg")
    img = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)

    img = resize(img, resize_shape)
    data = normalize(to_tensor(img), mean, std)
    data.unsqueeze_(0)

    seg_pred, exist_pred = net(data)[:2]
    seg_pred = F.softmax(seg_pred, dim=1)
    seg_pred = seg_pred.detach().cpu().numpy()
    exist_pred = exist_pred.detach().cpu().numpy()

    seg_pred = seg_pred[0]
    exist = [1 if exist_pred[0, i] > 0.5 else 0 for i in range(4)]

    img = cv2.cvtColor(img, cv2.COLOR_RGB2BGR)
    lane_img = np.zeros_like(img)

    color = np.array(
        [[255, 125, 0], [0, 255, 0], [0, 0, 255], [0, 255, 255]], dtype="uint8"
    )
    coord_mask = np.argmax(seg_pred, axis=0)
    for i in range(0, 4):
        if exist_pred[0, i] > 0.5:
            lane_img[coord_mask == (i + 1)] = color[i]
    img = cv2.addWeighted(src1=lane_img, alpha=0.8, src2=img, beta=1.0, gamma=0.0)
    cv2.imshow("", img)
    cv2.waitKey(0)
