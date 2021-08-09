TVM_ROOT=/home/sunway/source/tvm/
DMLC_CORE=${TVM_ROOT}/3rdparty/dmlc-core

ifeq ("$(wildcard $(TVM_ROOT))","")
$(error TVM_ROOT ${TVM_ROOT} doesnt exist)
endif

.PHONY: clean all run

APP=double

CPPFLAGS = -std=c++14 -O2 -fPIC \
		-I${TVM_ROOT}/include \
		-I${TVM_ROOT}/src \
		-I${DMLC_CORE}/include \
		-I${TVM_ROOT}/3rdparty/dlpack/include \
		-DDMLC_USE_LOGGING_LIBRARY=\<tvm/runtime/logging.h\> \
		-g \

CXXFLAGS = -Wvla -ffunction-sections -fdata-sections
LDFLAGS = -L${TVM_ROOT}/build -static -Wl,-Map=${APP}.map -Wl,-gc-sections

# comment out the following lines in library_module.cc to get rid of `-lpthread`
# TVM_INIT_CONTEXT_FUNC(TVMBackendParallelLaunch);
# TVM_INIT_CONTEXT_FUNC(TVMBackendParallelBarrier);
LDLIBS = -lstdc++ -lpthread

all: ${APP}

libdouble.o: libdouble.py
	python3 libdouble.py

OBJS=double.o libdouble.o  tvm_runtime.o
${APP}: ${OBJS}

clean:
	-rm ${OBJS}
	-rm ${APP}
	-rm ${APP}.map

run:${APP}
	${APP}
