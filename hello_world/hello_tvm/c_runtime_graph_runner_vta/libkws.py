#!/usr/bin/env python3
import os

import numpy as np
import tflite
import tvm
from tvm.contrib import graph_executor
from tvm import relay, runtime
import vta
from vta.top import graph_pack

from run_model import get_model

mod, params = get_model(mode="tvm_quant")
env = vta.get_env()

import ipdb

ipdb.set_trace()
mod = graph_pack(
    mod["main"],
    env.BATCH,
    env.BLOCK_OUT,
    env.WGT_WIDTH,
    start_name="nn.dense",
    stop_name="nn.dense",
    annot_start_name="multiply",
    annot_end_name="dense",
    device_annot=True,
)
with vta.build_config(opt_level=3, disabled_pass={"AlterOpLayout"}):
    graph, lib, params = relay.build(
        mod,
        target={"cpu": "llvm", "ext_dev": env.target},
        params=params,
        target_host="llvm",
    )


# target = "llvm  --system-lib --runtime=c"
# with tvm.transform.PassContext(opt_level=3):
#     mod = relay.build_module.build(mod, target=target, params=params)

# mod.lib.export_library("/tmp/libkws.tar")

# with open("kws_graph.json", "w") as f_graph_json:
#     f_graph_json.write(mod.graph_json)


# with open("kws_params.bin", "wb") as f_params:
#     f_params.write(runtime.save_param_dict(mod.params))
