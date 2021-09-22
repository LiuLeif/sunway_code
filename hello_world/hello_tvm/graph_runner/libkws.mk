TVM_ROOT=/home/sunway/source/tvm

DMLC_CORE=${TVM_ROOT}/3rdparty/dmlc-core
PKG_COMPILE_OPTS = -g -Wall -O2 -fPIC
CPPFLAGS = ${PKG_COMPILE_OPTS} \
	-I${TVM_ROOT}/include \
	-I${DMLC_CORE}/include \
	-I${TVM_ROOT}/3rdparty/dlpack/include \
	-I. \
	-DDMLC_USE_LOGGING_LIBRARY=\<tvm/runtime/logging.h\>

LDLIBS += -Wl,--whole-archive libkws.a -Wl,-no-whole-archive -lm

libkws.a:libkws.py
	python ./libkws.py --mode=${MODE}
	make -f /tmp/libkws/libkws.mk

common_clean:
	-rm -rf /tmp/libkws
	-rm libkws.a
