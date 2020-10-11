AS=mips-linux-gnu-as
LD=mips-linux-gnu-ld

all:test

%.o:%.mips
	${AS} $^ -o $@

test:test.o
	${LD} -G8 $^ -o $@

.PHONY: debug clean

run:test
	-pkill -9 qemu-mips-static
	qemu-mips-static -g 12345 ./test &
	gdb-multiarch -x ./gdb.cmds ./test

clean:
	-rm test.o test 2>/dev/null || true
