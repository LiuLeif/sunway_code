// 2022-06-08 14:47
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

extern void test1();
extern void test2();

int main(int argc, char *argv[]) {
    test1();
    test2();
}
