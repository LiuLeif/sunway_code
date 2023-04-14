all:native emu

CFLAGS := -O3 -g3
LDFLAGS := -L. -static
LDLIBS := -lm
CC :=aarch64-linux-gnu-gcc

SRC := $(wildcard test_*.c)
NATIVE_APP := $(patsubst %.c,native_%.elf,${SRC})
EMU_APP := $(patsubst %.c,emu_%.elf,${SRC})
TEST_NATIVE_APP := $(patsubst %.elf, run_%,${NATIVE_APP})
TEST_EMU_APP := $(patsubst %.elf, run_%,${EMU_APP})

native: ${NATIVE_APP}

emu: ${EMU_APP}

native_%.elf:%.c
	${CC} ${CFLAGS} -o $@ $^ ${LDFLAGS} ${LDLIBS}

emu_%.elf:%.c
	${CC} ${CFLAGS} -DEMU -o $@ $^ ${LDFLAGS} ${LDLIBS}

run_%:%.elf
	./$<

test_native: ${TEST_NATIVE_APP}
test_emu: ${TEST_EMU_APP}

run: test_native test_emu

clean:
	rm -f ${NATIVE_APP} ${EMU_APP}

.PHONY: clean app run
