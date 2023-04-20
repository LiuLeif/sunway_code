// 2023-04-20 17:45
#ifndef VCLS_H
#define VCLS_H

#include "neon_emu_types.h"
int8x8_t vcls_s8(int8x8_t a) {
    int8x8_t r = {0};
    for (int i = 0; i < 8; i++) {
        int8_t tmp = a.values[i];
        int8_t sign = tmp < 0 ? 1 : 0;
        for (int j = 0; i < 7; j++) {
            tmp <<= 1;
            if ((tmp & (1 << 7)) != sign) {
                break;
            }
            r.values[i] += 1;
        }
    }
    return r;
}

int8x8_t vcls_u8(uint8x8_t a) {
    int8x8_t r = {0};
    for (int i = 0; i < 8; i++) {
        uint8_t tmp = a.values[i];
        int8_t sign = (tmp & (1 << 7)) == 0 ? 0 : 1;
        for (int j = 0; i < 7; j++) {
            tmp <<= 1;
            if ((tmp & (1 << 7)) != sign) {
                break;
            }
            r.values[i] += 1;
        }
    }
    return r;
}

#endif  // VCLS_H
