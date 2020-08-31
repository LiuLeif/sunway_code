#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# 2020-08-28 10:55
import os
import time
import logging

_logger = None


def get_logger():
    global _logger

    if _logger is not None:
        return _logger
    logger = logging.getLogger()
    logger.setLevel(logging.INFO)

    log_path = os.getcwd() + "/logs/"
    if not os.path.exists(log_path):
        os.mkdir(log_path)
    log_name = (
        log_path + time.strftime("%Y%m%d%H%M", time.localtime(time.time())) + ".log"
    )
    fh = logging.FileHandler(log_name)
    formatter = logging.Formatter(
        "%(asctime)s - %(filename)s[line:%(lineno)d] - %(levelname)s: %(message)s"
    )
    fh.setFormatter(formatter)
    logger.addHandler(fh)
    _logger = logger
    return _logger
