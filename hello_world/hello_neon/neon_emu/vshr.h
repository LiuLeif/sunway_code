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
#endif  // VSHR_H
