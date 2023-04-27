// 2023-04-27 20:48
#include <stdint.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
    int8_t tmp = -1;
    int8_t sign = tmp & (1 << 7);
    if (sign == (tmp & (1 << 7))) {
        printf("ok\n");
    }
    return 0;
}
