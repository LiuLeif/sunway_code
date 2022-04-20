// 2022-04-20 21:26
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define DIV_ROUND_UP(m, n) (((m) + (n)-1) / (n))

static const char hex_digits[] = {'0', '1', '2', '3', '4', '5', '6', '7',
                                  '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'};

char *buf_to_hex_str(const void *_buf, unsigned buf_len) {
    unsigned len_bytes = DIV_ROUND_UP(buf_len, 8);
    char *str = calloc(len_bytes * 2 + 1, 1);

    const uint8_t *buf = _buf;
    for (unsigned i = 0; i < len_bytes; i++) {
        uint8_t tmp = buf[len_bytes - i - 1];
        if ((i == 0) && (buf_len % 8)) tmp &= (0xff >> (8 - (buf_len % 8)));
        str[2 * i] = hex_digits[tmp >> 4];
        str[2 * i + 1] = hex_digits[tmp & 0xf];
    }

    return str;
}

static int ceil_f_to_u32(float x) {
    if (x < 0) /* return zero for negative numbers */
        return 0;

    uint32_t y = x; /* cut off fraction */

    if ((x - y) > 0.0) /* if there was a fractional part, increase by one */
        y++;

    return y;
}

char *buf_to_str(const void *_buf, unsigned buf_len, unsigned radix) {
    float factor;
    switch (radix) {
        case 16:
            factor = 2.0; /* log(256) / log(16) = 2.0 */
            break;
        case 10:
            factor = 2.40824; /* log(256) / log(10) = 2.40824 */
            break;
        case 8:
            factor = 2.66667; /* log(256) / log(8) = 2.66667 */
            break;
        default:
            return NULL;
    }

    unsigned str_len = ceil_f_to_u32(DIV_ROUND_UP(buf_len, 8) * factor);
    char *str = calloc(str_len + 1, 1);

    const uint8_t *buf = _buf;
    int b256_len = DIV_ROUND_UP(buf_len, 8);
    for (int i = b256_len - 1; i >= 0; i--) {
        uint32_t tmp = buf[i];
        if (((unsigned)i == (buf_len / 8)) && (buf_len % 8))
            tmp &= (0xff >> (8 - (buf_len % 8)));

        /* base-256 digits */
        for (unsigned j = str_len; j > 0; j--) {
            tmp += (uint32_t)str[j - 1] * 256;
            str[j - 1] = (uint8_t)(tmp % radix);
            tmp /= radix;
        }
    }

    const char *const DIGITS = "0123456789ABCDEF";
    for (unsigned j = 0; j < str_len; j++) str[j] = DIGITS[(int)str[j]];

    return str;
}

int main(int argc, char *argv[]) {
    char buffer[] = {0xab, 0xcd, 0xef};
    printf("%s\n", buf_to_hex_str((void *)buffer, 24));
    printf("%s\n", buf_to_str((void *)buffer, 24, 10));
}
