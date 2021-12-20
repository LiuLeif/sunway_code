# 1. Get ComputeCPP from https://developer.codeplay.com/products/computecpp/ce/download/
# 2. apt install intel-opencl-icd
SRC := $(shell find ./ -name "*.cc")
CPPFLAGS := -I ~/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/include/
LDFLAGS := -L ~/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/lib/ -Wl,-rpath,/home/sunway/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/lib/
LDLIBS := -lComputeCpp -lOpenCL

.PHONY: ALL clean

APP := $(patsubst %.cc,%.elf,${SRC})

ALL: ${APP}
# %.elf:%.cc
# 	~/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/bin/compute $^ ${CPPFLAGS} ${LDFLAGS} ${LDLIBS} -sycl-driver -o $@

%.sycl:%.cc
	~/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/bin/compute $^ ${CPPFLAGS} -sycl -c $^

%.elf: %.sycl %.cc
	g++ ${CPPFLAGS} ${LDFLAGS} -include $^ -o $@ ${LDLIBS}

clean:
	-rm -f *.elf
	-rm -f *.sycl
