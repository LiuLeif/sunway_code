
all: test

GCCPLUGINS_DIR := $(shell gcc -print-file-name=plugin)
PLUGIN_CXXFLAGS := -I$(GCCPLUGINS_DIR)/include -fPIC -fno-rtti -O2

PLUGIN_SRC := $(wildcard *.cc)
PLUGIN_OBJ := $(patsubst %.cc,%.o,${PLUGIN_SRC})

$(PLUGIN_OBJ):CXXFLAGS := ${PLUGIN_CXXFLAGS}

PLUGIN := libhello.so
${PLUGIN}: $(PLUGIN_OBJ)
	gcc -shared $^ -o $@ -lstdc++

#-------------------------------------------
TEST_CFLAGS := -O0 -g3 -fplugin=./${PLUGIN} -fplugin-arg-libhello-count=1

TEST_SRC := $(wildcard test*.c)
TEST_OBJ := $(patsubst %.c,%.o,${TEST_SRC})
${TEST_OBJ}:CFLAGS:=${TEST_CFLAGS}
${TEST_OBJ}:${PLUGIN} FORCE

test:${TEST_OBJ}

FORCE:

clean:
	rm -f ${TEST_OBJ} ${PLUGIN} ${PLUGIN_OBJ}

.PHONY: test clean FORCE
