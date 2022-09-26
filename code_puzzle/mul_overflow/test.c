// 2022-09-26 14:47
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int main(int argc, char *argv[]) {
    int32_t a = 0x40000000;
    int32_t b = 0x4;
    /* NOTE: 两个 int32_t 相乘结果是 int32_t */
    int64_t c = a * b;
    int64_t d = (int64_t)a * b;
    printf("%ld %ld\n", c, d);
}
