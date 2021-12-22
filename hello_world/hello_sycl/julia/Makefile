ALL:julia.elf

.PHONY: ALL clean

SYCL_CPPFLAGS := -I ~/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/include/
SYCL_LDFLAGS := -L ~/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/lib/ -Wl,-rpath,/home/sunway/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/lib/
SYCL_LDLIBS := -lComputeCpp -lOpenCL

CPPFLAGS := -I/home/sunway/download/Cinder/include -MMD
LDLIBS := /home/sunway/download/Cinder/lib/linux/x86_64/ogl/Debug/libcinder.a \
		  -ldl -lOpenGL -lcurl -lpthread -lX11 -lfontconfig -lXrandr -lXinerama -lXcursor

CPPFLAGS += ${SYCL_CPPFLAGS}
LDFLAGS += ${SYCL_LDFLAGS}
LDLIBS += ${SYCL_LDLIBS}

SRC:=$(wildcard *.cc)
OBJ:=$(patsubst %.cc,%.o,${SRC})

DEPS := $(OBJ:.o=.d)
-include $(DEPS)

SYCL_SRC:=$(wildcard *.sycl.cc)
SYCL_HEADER=$(patsubst %.cc,%.sycl,${SYCL_SRC})

%.sycl:%.cc
	~/program/ComputeCpp-CE-2.8.0-x86_64-linux-gnu/bin/compute $^ ${SYCL_CPPFLAGS} -sycl -c $^

julia.elf:${OBJ} ${SYCL_HEADER}
	g++ ${OBJ} -o $@ ${LDFLAGS} ${LDLIBS} -include ${SYCL_HEADER}

clean:
	-rm -f julia.elf
	-rm -f ${OBJ}
	-rm -f ${DEPS}
