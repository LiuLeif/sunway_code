// 2022-06-08 14:47
#include <stdio.h>

extern int xxx[10];
void test2() {
    for (int i = 0; i < 10; i++) {
        xxx[i] = 1;
    }
    printf("test2\n");
};
