all: hello_fft

LDLIBS := kissfft/libkissfft-float.a -lm

hello_fft:kissfft/libkissfft-float.a

kissfft/libkissfft-float.a: FORCE
	make -C kissfft KISSFFT_STATIC=1

FORCE:

clean:
	rm hello_fft

run:hello_fft
	./hello_fft
