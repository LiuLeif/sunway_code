// 2023-02-01 11:05
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]) {
    int32_t x = 0x40000000;
    int32_t y = 0x70000000;
    int64_t a = (int64_t)(x * y / 2147483648.0);
    int64_t b = (int64_t)(x * 1.0 * y / 2147483648.0);
    printf("%ld %ld\n", a, b);
    return 0;
}
