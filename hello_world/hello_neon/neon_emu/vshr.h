// 2023-04-19 16:24
#ifndef VSHR_H
#define VSHR_H

#include "neon_emu_types.h"

int16x4_t vshr_n_s16(int16x4_t a, int n) {
    int16x4_t r;
    for (int i = 0; i < 4; i++) {
        r.values[i] = a.values[i] >> n;
    }
    return r;
}

uint16x4_t vshr_n_u16(uint16x4_t a, int n) {
    uint16x4_t r;
    for (int i = 0; i < 4; i++) {
        r.values[i] = a.values[i] >> n;
    }
    return r;
}

int8x8_t vrshr_n_s8(int8x8_t a, int n) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = (a.values[i] + (1 << (n - 1))) >> n;
    }
    return r;
}

int8x8_t vsra_n_s8(int8x8_t a, int8x8_t b, int n) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] + (b.values[i] >> n);
    }
    return r;
}

int8x8_t vrsra_n_s8(int8x8_t a, int8x8_t b, int n) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] + ((b.values[i] + (1 << (n - 1))) >> n);
    }
    return r;
}
#endif  // VSHR_H
