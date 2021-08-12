#!/usr/bin/env python3
import os

import numpy as np
import tflite
import tvm
from tvm.contrib import graph_executor
from tvm import relay, runtime

tflite_model = tflite.Model.GetRootAsModel(open("kws.tflite", "rb").read(), 0)

mod, params = relay.frontend.from_tflite(
    tflite_model, shape_dict={"input_1": (1, 99, 12)}, dtype_dict={"input_1": "float32"}
)

target = "llvm  --system-lib --runtime=c"
with tvm.transform.PassContext(opt_level=3):
    mod = relay.build_module.build(mod, target=target, params=params)

mod.lib.export_library("/tmp/libkws.tar")

with open("kws_graph.json", "w") as f_graph_json:
    f_graph_json.write(mod.graph_json)

with open("kws_params.bin", "wb") as f_params:
    f_params.write(runtime.save_param_dict(mod.params))
