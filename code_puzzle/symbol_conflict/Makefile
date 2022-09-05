main: main.o libtest1.a libtest2.a
CFLAGS := -O0 -g
LDFLAGS :=-Wl,-Map=test.map
libtest1.a:test1.o test11.o
	ar rcs $@ $^
libtest2.a:test2.o test22.o
	ar rcs $@ $^
clean:
	rm *.o *.a main
