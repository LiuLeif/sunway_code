// 2023-04-27 20:48
#include <stdint.h>
#include <stdio.h>

int main(int argc, char *argv[]) {
    int8_t tmp = -1;
    /* NOTE: sign 为 0x80 */
    int8_t sign = tmp & (1 << 7);
    /* NOTE: gcc 先把左侧 sign 符号扩展成 int32_t, 为 0xffffff80, 右侧也是以
     * int32_t 为计算, 导致结果为 0x00000080. 如果 int8_t 换成 uint8_t 则没有问
     * 题, 因为没有符号扩展. 所以 bit 操作最好用 unsigned */
    if (sign == (tmp & (1 << 7))) {
        printf("ok\n");
    }
    return 0;
}
