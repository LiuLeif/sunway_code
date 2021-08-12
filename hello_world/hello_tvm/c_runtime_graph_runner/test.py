#!/usr/bin/env python3
import os

import numpy as np
import tflite
import tvm
from tvm.contrib import graph_executor
from tvm import relay

tflite_model = tflite.Model.GetRootAsModel(open("kws.tflite", "rb").read(), 0)

mod, params = relay.frontend.from_tflite(
    tflite_model, shape_dict={"input_1": (1, 99, 12)}, dtype_dict={"input_1": "float32"}
)

target = "llvm"
with tvm.transform.PassContext(opt_level=3):
    lib = relay.build_module.build(mod, target=target, params=params)

rt_mod = graph_executor.GraphModule(lib["default"](tvm.cpu(0)))
x = np.load("test/test_xiaoai.npy")
for i in range(len(x)):
    rt_mod.set_input("input_1", x[i : i + 1])
    rt_mod.run()
    tvm_res = rt_mod.get_output(0).numpy()
    print(tvm_res)

x = np.load("test/test_unknown.npy")
for i in range(len(x)):
    rt_mod.set_input("input_1", x[i : i + 1])
    rt_mod.run()
    tvm_res = rt_mod.get_output(0).numpy()
    print(tvm_res)
