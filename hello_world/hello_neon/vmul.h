// 2023-04-14 19:39
#ifndef VMUL_H
#define VMUL_H
#include "neon_emu_types.h"

int8x8_t vmul_s8(int8x8_t a, int8x8_t b) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] * b.values[i];
    }
    return r;
}

int8x16_t vmulq_s8(int8x16_t a, int8x16_t b) {
    int8x16_t r;
    for (int i = 0; i < 16; i++) {
        r.values[i] = a.values[i] * b.values[i];
    }
    return r;
}

#endif  // VMUL_H
