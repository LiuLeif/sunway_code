TEST:=$(patsubst %.c,%.elf,$(wildcard test_*.c))

LDLIBS:=-lbfd
CFLAGS:=-g

all:${TEST}

%.elf:%.c
	${CC} ${CFLAGS} -o $@ $< ${LDLIBS}

clean:
	-rm -rf ${TEST}
