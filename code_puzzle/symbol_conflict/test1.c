// 2022-06-08 14:47
#include <stdio.h>

extern int xxx[200000];
void test1() {
    for (int i = 0; i < 200000; i++) {
        xxx[i] = 1;
    }
    printf("test1\n");
}
