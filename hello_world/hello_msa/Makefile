all:native emu

CFLAGS := -mmsa -MMD -O3 -g3 -Imsa_emu -I.
LDFLAGS := -L. -static
LDLIBS := -lm
CC := ~/program/toolchain/mips-gcc720-glibc229/bin/mips-linux-gnu-gcc

SRC := $(shell find ./ -name test_*.c -printf '%P\n')
NATIVE_APP := $(patsubst %.c,%_native.elf,${SRC})
EMU_APP := $(patsubst %.c,%_emu.elf,${SRC})
TEST_NATIVE_APP := $(patsubst %, run_%,${NATIVE_APP})
TEST_EMU_APP := $(patsubst %, run_%,${EMU_APP})

native: ${NATIVE_APP}
emu: ${EMU_APP}

DEP := $(NATIVE_APP:.elf=.d)
DEP += $(EMU_APP:.elf=.d)
-include ${DEP}

%_native.elf:%.c
	${CC} ${CFLAGS} -o $@ $< ${LDFLAGS} ${LDLIBS}
%_emu.elf:%.c
	${CC} ${CFLAGS} -DEMU -o $@ $< ${LDFLAGS} ${LDLIBS}

${TEST_EMU_APP}:run_%:%
	qemu-mipsel -cpu P5600 ./$<

${TEST_NATIVE_APP}:run_%:%
	qemu-mipsel -cpu P5600 ./$<

run_native: ${TEST_NATIVE_APP}
run_emu: ${TEST_EMU_APP}
run: run_native run_emu

clean:
	rm -f ${NATIVE_APP} ${EMU_APP} ${DEP}

.PHONY: clean app run
