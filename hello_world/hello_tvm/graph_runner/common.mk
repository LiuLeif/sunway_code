TVM_ROOT=/home/sunway/source/tvm

DMLC_CORE=${TVM_ROOT}/3rdparty/dmlc-core
PKG_COMPILE_OPTS = -g -Wall -O2 -fPIC
CPPFLAGS = ${PKG_COMPILE_OPTS} \
	-I${TVM_ROOT}/include \
	-I${DMLC_CORE}/include \
	-I${TVM_ROOT}/3rdparty/dlpack/include \
	-I. \
	-DDMLC_USE_LOGGING_LIBRARY=\<tvm/runtime/logging.h\> \
	# -ffunction-sections -fdata-sections

# LDFLAGS = -static # -Wl,-gc-sections
LDLIBS = -lm

ifeq (${MODE},c)
	MODEL_OBJ = devc.o lib0.o lib1.o
endif

ifeq (${MODE},dnnl)
	MODEL_OBJ = lib0.o devc.o
endif

ifeq (${MODE},c++)
	MODEL_OBJ = lib0.o
endif

MODEL_OBJ_AUX = kws_graph.o kws_params.o

kws_graph.c:kws_graph.json
	xxd -i $^ > $@

kws_params.c:kws_params.bin
	xxd -i $^ > $@


${MODEL_OBJ} kws_graph.json kws_params.bin:libkws.py
	python ./libkws.py --mode=${MODE}
	tar xvf /tmp/libkws.tar

common_clean:
	-rm ${MODEL_OBJ}
	-rm ${MODEL_OBJ_AUX}
	-rm kws_graph.c
	-rm kws_params.c
	-rm kws_graph.json
	-rm kws_params.bin
