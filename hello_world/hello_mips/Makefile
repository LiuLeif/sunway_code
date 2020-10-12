AS=mips-linux-gnu-as
LD=mips-linux-gnu-ld

all:test

%.o:%.mips
	${AS} $^ -o $@ -gstabs

test:test.o lib.o
	${LD} -G8 $^ -o $@

.PHONY: debug clean

gdb:test
	-pkill -9 qemu-mips-static || true
	cat test.stdin|qemu-mips-static -g 12345 ./test &

run:test
	-pkill -9 qemu-mips-static || true
	cat test.stdin|qemu-mips-static ./test

clean:
	-rm test.o test 2>/dev/null || true
