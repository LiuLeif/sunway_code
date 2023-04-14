all:native emu

CFLAGS := -O3 -g3 -I.
LDFLAGS := -L. -static
LDLIBS := -lm
CC :=aarch64-linux-gnu-gcc

SRC := $(shell find ./ -name test_*.c)
NATIVE_APP := $(patsubst %.c,%_native.elf,${SRC})
EMU_APP := $(patsubst %.c,%_emu.elf,${SRC})
TEST_NATIVE_APP := $(patsubst %, run_%,${NATIVE_APP})
TEST_EMU_APP := $(patsubst %, run_%,${EMU_APP})

native: ${NATIVE_APP}
emu: ${EMU_APP}

%_native.elf:%.c
	${CC} ${CFLAGS} -o $@ $^ ${LDFLAGS} ${LDLIBS}
%_emu.elf:%.c
	${CC} ${CFLAGS} -DEMU -o $@ $^ ${LDFLAGS} ${LDLIBS}

${TEST_EMU_APP}:run_%:%
	./$<

${TEST_NATIVE_APP}:run_%:%
	./$<

run_native: ${TEST_NATIVE_APP}
run_emu: ${TEST_EMU_APP}
run: run_native run_emu

clean:
	rm -f ${NATIVE_APP} ${EMU_APP}

.PHONY: clean app run
