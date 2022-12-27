CFLAGS := -I/usr/lib/x86_64-linux-gnu/openmpi/include/
LDFLAGS := -lmpi
SRC := $(wildcard *.c)
APP := $(patsubst %.c, %.elf,${SRC})
RUN_APP := $(patsubst %.elf, run_%,${APP})
# CC := mpicc

all:${APP}

%.elf: %.c
	${CC} -o $@ $^ ${CFLAGS} ${LDFLAGS}

${RUN_APP}:run_%:%.elf
	mpirun -np 4 $<

clean:
	-rm -f ${APP}
