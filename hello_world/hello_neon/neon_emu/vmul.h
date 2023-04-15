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

int8x8_t vmla_s8(int8x8_t a, int8x8_t b, int8x8_t c) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] + b.values[i] * c.values[i];
    }
    return r;
}

int8x8_t vmls_s8(int8x8_t a, int8x8_t b, int8x8_t c) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] - b.values[i] * c.values[i];
    }
    return r;
}

int16x8_t vmlal_s8(int16x8_t a, int8x8_t b, int8x8_t c) {
    int16x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] + b.values[i] * c.values[i];
    }
    return r;
}

int16x8_t vmlsl_s8(int16x8_t a, int8x8_t b, int8x8_t c) {
    int16x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] - b.values[i] * c.values[i];
    }
    return r;
}

int16x8_t vmlal_high_s8(int16x8_t a, int8x16_t b, int8x16_t c) {
    int16x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] + b.values[i + 8] * c.values[i + 8];
    }
    return r;
}

int16x8_t vmlsl_high_s8(int16x8_t a, int8x16_t b, int8x16_t c) {
    int16x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] - b.values[i + 8] * c.values[i + 8];
    }
    return r;
}

#endif  // VMUL_H
