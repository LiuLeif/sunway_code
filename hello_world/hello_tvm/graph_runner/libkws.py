#!/usr/bin/env python3
import os
import argparse

import numpy as np
import tflite
import tvm
from tvm.contrib import graph_executor
from tvm import relay, runtime

from run_model import get_model

parser = argparse.ArgumentParser()
parser.add_argument("--mode", choices=["c", "c++", "dnnl"], required=True)
args = parser.parse_args()

mod, params = get_model(mode="float")

if args.mode in ("c", "c++"):
    target = f"llvm  --system-lib --runtime={args.mode}"

if args.mode == "dnnl":
    # due to dnnl byoc bugs: these following lines need to be commented out from dnnl.py
    # _register_external_op_helper("nn.relu")
    # _register_external_op_helper("add")
    # _register_external_op_helper("subtract")
    # _register_external_op_helper("multiply")
    seq = tvm.transform.Sequential(
        [relay.transform.ConvertLayout({"nn.conv2d": ["NCHW", "default"]})]
    )
    with tvm.transform.PassContext(opt_level=3):
        mod = seq(mod)
    mod = relay.transform.AnnotateTarget("dnnl")(mod)
    mod = relay.transform.PartitionGraph()(mod)
    target = f"llvm  --system-lib --runtime=c++"

with tvm.transform.PassContext(opt_level=3):
    mod = relay.build_module.build(mod, target=target, params=params)

mod.lib.export_library("/tmp/libkws.tar")

with open("kws_graph.json", "w") as f_graph_json:
    f_graph_json.write(mod.graph_json)


with open("kws_params.bin", "wb") as f_params:
    f_params.write(runtime.save_param_dict(mod.params))
