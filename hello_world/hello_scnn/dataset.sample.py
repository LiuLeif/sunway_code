#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2021-12-30 19:14
import json
import os

import cv2
import numpy as np
import torch
from torch.utils.data import Dataset

# NOTE: 这是简化版本的 dataset, 删减了许多东西, 代码并不能直接运行
class Tusimple(Dataset):
    def __init__(self, path, image_set, transforms=None):
        super(Tusimple, self).__init__()
        self.data_dir_path = path
        self.image_set = image_set

        if not os.path.exists(os.path.join(path, "seg_label")):
            # NOTE: generate_label 主要是根据描述坐标的 json 文件 (例如
            # hello_tusimple.json) 生成 prob_map 格式的文件, 后续会直接用
            # cross_entropy 来计算 seg_pred 与 seg_gt 的 loss
            self.generate_label()
        self.createIndex()

    def createIndex(self):
        listfile = os.path.join(
            self.data_dir_path, "seg_label", "list", "{}_gt.txt".format(self.image_set)
        )
        with open(listfile) as f:
            for line in f:
                line = line.strip()
                l = line.split(" ")
                # segLabel_list 格式是 [xxx.png, yyy.png, zzz.png ...]
                # exist_list 格式是 [[0 0 0 0], [1,1,0,0],[1,0,1,1],...]
                self.segLabel_list.append(os.path.join(self.data_dir_path, l[1][1:]))
                self.exist_list.append([int(x) for x in l[2:]])

    def __getitem__(self, idx):
        img = cv2.imread(self.img_list[idx])
        img = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)

        # NOTE: imread 返回的是一个 BGR 的图片, 但这里只需要 b,g,r 的某一个来代
        # 表像素所属的分类, 因为 bgr 在 generate_label 时写的是值是相同的
        segLabel = cv2.imread(self.segLabel_list[idx])[:, :, 0]
        exist = np.array(self.exist_list[idx])

        sample = {
            "img": img,
            "segLabel": segLabel,
            "exist": exist,
            "img_name": self.img_list[idx],
        }
        return sample

    def generate_label(self):
        self._gen_label_for_json("train")
        print("train set is done")
        self._gen_label_for_json("val")
        print("val set is done")

    def _gen_label_for_json(self, image_set):
        # NOTE: tusimple 图片的尺寸是 1280x720
        H, W = 720, 1280

        # NOTE: SEG_WIDTH 是指 cv2.line 的宽度...
        SEG_WIDTH = 30

        with open(json_path) as f:
            for line in f:
                # NOTE: seg_img 初始为 0, 代码像素属于 backgroud
                seg_img = np.zeros((H, W, 3))
                for i in range(len(lanes)):
                    coords = lanes[i]
                    if len(coords) < 4:
                        # NOTE: 当某条 lane 的点很少, 则记录 exist_list 为 0
                        list_str.append("0")
                        continue
                    for j in range(len(coords) - 1):
                        # NOTE: 这里用了一个 trick: 把所有的点连接起来, 同时指定 color 为 (i+1),
                        # 这样生成的 png 文件中每个像素的 color 即是它的车道线分类
                        cv2.line(
                            seg_img,
                            coords[j],
                            coords[j + 1],
                            (i + 1, i + 1, i + 1),
                            SEG_WIDTH // 2,
                        )
                    list_str.append("1")

                seg_path = os.path.join(seg_path, img_name[:-3] + "png")
                cv2.imwrite(seg_path, seg_img)

                seg_path = "/".join(
                    [save_dir, *img_path.split("/")[1:3], img_name[:-3] + "png"]
                )
                list_str.insert(0, seg_path)
                list_str.insert(0, img_path)
                list_str = " ".join(list_str) + "\n"
                list_f.write(list_str)

        list_f.close()
