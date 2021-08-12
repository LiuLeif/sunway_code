all: kws

TVM_ROOT=/home/sunway/source/tvm
CRT_ROOT=${TVM_ROOT}/build/standalone_crt

DMLC_CORE=${TVM_ROOT}/3rdparty/dmlc-core
PKG_COMPILE_OPTS = -g -Wall -O2 -fPIC
CFLAGS = ${PKG_COMPILE_OPTS} \
	-I${TVM_ROOT}/include \
	-I${DMLC_CORE}/include \
	-I${TVM_ROOT}/3rdparty/dlpack/include \
	-I. \
	-DDMLC_USE_LOGGING_LIBRARY=\<tvm/runtime/logging.h\> \
	-ffunction-sections -fdata-sections

CFLAGS += -g
LDFLAGS = -static -Wl,-gc-sections
LDLIBS = -lm

CRT_SRCS = $(shell find $(CRT_ROOT))

MODEL_OBJ = devc.o lib0.o lib1.o

kws_graph.c:kws_graph.json
	xxd -i $^ > $@

kws_params.c:kws_params.bin
	xxd -i $^ > $@

${MODEL_OBJ} kws_graph.json kws_params.bin:libkws.py
	python ./libkws.py
	tar xvf /tmp/libkws.tar

crt/libcommon.a: $(CRT_SRCS)
	cd $(CRT_ROOT) && make QUIET= BUILD_DIR=$(abspath .)/crt CRT_CONFIG=$(abspath crt_config.h) "EXTRA_CFLAGS=$(PKG_COMPILE_OPTS)" common

crt/libmemory.a: $(CRT_SRCS)
	cd $(CRT_ROOT) && make QUIET= BUILD_DIR=$(abspath .)/crt CRT_CONFIG=$(abspath crt_config.h) "EXTRA_CFLAGS=$(PKG_COMPILE_OPTS)" memory

crt/libgraph_executor.a: $(CRT_SRCS)
	$(QUIET)cd $(CRT_ROOT) && make QUIET= BUILD_DIR=$(abspath .)/crt CRT_CONFIG=$(abspath crt_config.h) "EXTRA_CFLAGS=$(PKG_COMPILE_OPTS)" graph_executor

OBJS=kws.o tvm_runtime.o kws_graph.o kws_params.o test_data.o

kws: ${OBJS} $(MODEL_OBJ) crt/libmemory.a crt/libcommon.a crt/libgraph_executor.a

run:kws
	./kws

clean:
	-rm ${OBJS}
	-rm ${MODEL_OBJ}
	-rm kws
	-rm -rf crt
	-rm kws_graph.json
	-rm kws_params.bin
