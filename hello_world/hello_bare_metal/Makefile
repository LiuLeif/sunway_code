OBJ=startup.o test.o
CFLAGS= -mcpu=arm926ej-s
ASFLAGS = -mcpu=arm926ej-s

CC=arm-none-eabi-gcc -g
AS=arm-none-eabi-as
LD=arm-none-eabi-ld
OBJCOPY=arm-none-eabi-objcopy

all: test.bin

test.elf:${OBJ}
	${LD} -T test.ld $^ -o $@

test.bin:test.elf
	${OBJCOPY} -O binary $^ $@

clean:
	-rm ${OBJ}
	-rm test.bin
	-rm test.elf

run:
	qemu-system-arm -M versatilepb -m 128M -nographic -kernel test.bin
