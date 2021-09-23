#!/usr/bin/env python3
import os
import argparse

import numpy as np
import tflite
import tvm
from tvm.contrib import graph_executor
from tvm import relay, runtime
import shutil
import tarfile

from run_model import get_model

parser = argparse.ArgumentParser()
parser.add_argument("--mode", choices=["c", "c++", "dnnl"], required=True)
args = parser.parse_args()

if args.mode in ("c", "c++"):
    mod, params = get_model(mode="tvm_quant")
    target = f"llvm  --system-lib --runtime={args.mode}"

if args.mode == "dnnl":
    mod, params = get_model(mode="float")
    target = f"llvm  --system-lib --runtime=c"
    # NOTE: dnnl byoc implementation sucks, see ./dnnl.patch for details
    seq = tvm.transform.Sequential(
        [relay.transform.ConvertLayout({"nn.conv2d": ["NCHW", "OIHW"]})]
    )
    with tvm.transform.PassContext(opt_level=3):
        mod = seq(mod)

    mod = relay.transform.AnnotateTarget("dnnl")(mod)
    mod = relay.transform.PartitionGraph()(mod)

with tvm.transform.PassContext(opt_level=3):
    mod = relay.build_module.build(mod, target=target, params=params)

mod.lib.export_library("/tmp/libkws.tar", cc = "c++")

shutil.rmtree("/tmp/libkws", ignore_errors = True)
tarfile.open("/tmp/libkws.tar").extractall("/tmp/libkws")
with open("/tmp/libkws/kws_graph.json", "w") as f_graph_json:
    f_graph_json.write(mod.graph_json)

with open("/tmp/libkws/kws_params.bin", "wb") as f_params:
    f_params.write(runtime.save_param_dict(mod.params))

os.system("cd /tmp/libkws/ && xxd -i kws_params.bin > kws_params.c")
os.system("cd /tmp/libkws/ && xxd -i kws_graph.json > kws_graph.c")
with open("/tmp/libkws/libkws.mk", "w") as f:
    f.write(
        """
TVM_ROOT=/home/sunway/source/tvm

DMLC_CORE=${TVM_ROOT}/3rdparty/dmlc-core
PKG_COMPILE_OPTS = -g -Wall -O2 -fPIC
CPPFLAGS = ${PKG_COMPILE_OPTS} \
	-I${TVM_ROOT}/include \
	-I${TVM_ROOT}/src/runtime/contrib/ \
	-I${DMLC_CORE}/include \
	-I${TVM_ROOT}/3rdparty/dlpack/include \
	-I. \
	-DDMLC_USE_LOGGING_LIBRARY=\<tvm/runtime/logging.h\>

libkws.a:$(wildcard /tmp/libkws/*.o)
libkws.a:$(patsubst %.cc,%.o,$(wildcard /tmp/libkws/*.cc))
libkws.a:
	ar rcs $@ $^
"""
    )
