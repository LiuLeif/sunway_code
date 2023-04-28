all:native c msa

CFLAGS := -MMD -O0 -g3 -I.
LDFLAGS := -L. -static
LDLIBS := -lm
CC :=aarch64-linux-gnu-gcc

SRC := $(shell find ./ -name test_*.c -printf '%P\n')
NATIVE_APP := $(patsubst %.c,%_native.elf,${SRC})
C_APP := $(patsubst %.c,%_c.elf,${SRC})
MSA_APP := $(patsubst %.c,%_msa.elf,${SRC})

TEST_NATIVE_APP := $(patsubst %, run_%,${NATIVE_APP})
TEST_C_APP := $(patsubst %, run_%,${C_APP})
TEST_MSA_APP := $(patsubst %, run_%,${MSA_APP})

native: ${NATIVE_APP}
c: ${C_APP}
msa: ${MSA_APP}

DEP := $(NATIVE_APP:.elf=.d)
DEP += $(C_APP:.elf=.d)
-include ${DEP}

%_native.elf:%.c
	aarch64-linux-gnu-gcc ${CFLAGS} -o $@ $< ${LDFLAGS} ${LDLIBS}
%_c.elf:%.c
	gcc ${CFLAGS} -Ineon_c -o $@ $< ${LDFLAGS} ${LDLIBS}
%_msa.elf:%.c
	mips64-linux-gnuabi64-gcc -mmsa -march=loongson3a -Ineon_msa ${CFLAGS}  -o $@ $< ${LDFLAGS} ${LDLIBS}


${TEST_C_APP}:run_%:%
	./$<

${TEST_NATIVE_APP}:run_%:%
	qemu-aarch64 ./$<

${TEST_MSA_APP}:run_%:%
	/opt/qemu/bin/qemu-mips64 --cpu Loongson-3A4000 ./$<

run_native: ${TEST_NATIVE_APP}
run_c: ${TEST_C_APP}
run_msa: ${TEST_MSA_APP}
run: run_native run_c run_msa

clean:
	rm -f ${NATIVE_APP} ${C_APP} ${DEP}

.PHONY: clean app run
