SRCS := $(wildcard *.cpp)
BINS := $(SRCS:%.cpp=%)

CFLAGS=-I ACADOtoolkit/ -I ACADOtoolkit/acado/ -g -O2

.PHONY: all
all: $(BINS)

$(BINS):%:%.cpp
	clang++ $<  $(CFLAGS) ACADOtoolkit/build/lib/libacado_toolkit_s.so -Wl,-rpath ACADOtoolkit/build/lib/ -o $@

clean:
	rm -rf $(BINS)
