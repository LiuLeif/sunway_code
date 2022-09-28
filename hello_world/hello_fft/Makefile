all: hello_fft

LDLIBS := kissfft/libkissfft-float.a -lm

CFLAGS := -g -Og

hello_fft:kissfft/libkissfft-float.a hello_fft.o twi_table.o

kissfft/libkissfft-float.a: FORCE
	make -C kissfft KISSFFT_STATIC=1

FORCE:

clean:
	rm hello_fft *.o

run:hello_fft
	./hello_fft
