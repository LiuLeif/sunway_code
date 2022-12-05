CFLAGS := -fopenmp
SRC := $(wildcard *.c)
APP := $(patsubst %.c, %.elf,${SRC})

all:${APP}

%.elf: %.c
	gcc -o $@ $^ ${CFLAGS}

clean:
	-rm -f ${APP}
