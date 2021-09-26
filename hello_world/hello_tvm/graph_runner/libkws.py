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
from tvm.relay.op.contrib import get_pattern_table

from run_model import get_model

parser = argparse.ArgumentParser()
parser.add_argument("--runtime", choices=["c", "c++"], required=True)
parser.add_argument("--dnnl", action="store_true")
args = parser.parse_args()

target = f"llvm  --system-lib --runtime={args.runtime}"

if args.dnnl:
    mod, params = get_model(mode="float")
    print(mod)
    # NOTE: dnnl byoc implementation sucks, see ./README.org for details
    dnnl_patterns = get_pattern_table("dnnl")

    seq = tvm.transform.Sequential(
        [
            relay.transform.MergeComposite(dnnl_patterns),
            relay.transform.AnnotateTarget("dnnl"),
            relay.transform.PartitionGraph(),
            relay.transform.ConvertLayout({"nn.conv2d": ["NCHW", "OIHW"]}),
        ]
    )
    with tvm.transform.PassContext(opt_level=3):
        mod = seq(mod)

else:
    mod, params = get_model(mode="float")
print(mod)
with tvm.transform.PassContext(opt_level=2):
    mod = relay.build_module.build(mod, target=target, params=params)

mod.lib.export_library("/tmp/libkws.tar")

shutil.rmtree("/tmp/libkws", ignore_errors=True)
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
libkws.a:$(wildcard /tmp/libkws/*.o)
libkws.a:$(patsubst %.c,%.o,$(wildcard /tmp/libkws/*.c))
%.o:%.c
	${CXX} ${CPPFLAGS} $< -c -o $@
libkws.a:
	ar rcs $@ $^
"""
    )
