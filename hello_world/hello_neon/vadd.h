// 2023-04-14 12:20
#ifndef VADD_H
#define VADD_H
#include <stdint.h>

#include "neon_emu_types.h"
// NOTE: 使用 64-bit vector 计算 a+b
int8x8_t vadd_s8(int8x8_t a, int8x8_t b) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] + b.values[i];
    }
    return r;
}
// NOTE: 使用 128-bit vector 计算 a+b, q 后缀表示 128-bit
int8x16_t vaddq_s8(int8x16_t a, int8x16_t b) {
    int8x16_t r;
    for (int i = 0; i < 16; i++) {
        r.values[i] = a.values[i] + b.values[i];
    }
    return r;
}

// NOTE: 返回 a+b 的高 8 位, hn 后缀表示 high narrow
// 不存在 vaddhn_s8, 因为 s8 无法再 narrow
int8x8_t vaddhn_s16(int16x8_t a, int16x8_t b) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = (a.values[i] + b.values[i]) >> 8;
    }
    return r;
}

// NOTE: 返回 (a+b)/2, h 前缀表示 half
int8x8_t vhadd_s8(int8x8_t a, int8x8_t b) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = (a.values[i] + b.values[i]) >> 1;
    }
    return r;
}

// NOTE: 返回 round((a+b)/2), rh 前缀表示 round half
int8x8_t vrhadd_s8(int8x8_t a, int8x8_t b) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = (a.values[i] + b.values[i] + 1) >> 1;
    }
    return r;
}

// NOTE: 返回 saturate(a+b), q 前缀表示 saturate
int8x8_t vqadd_s8(int8x8_t a, int8x8_t b) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        int16_t tmp = (int16_t)a.values[i] + b.values[i];
        if (tmp > INT8_MAX) {
            tmp = INT8_MAX;
        }
        if (tmp < INT8_MIN) {
            tmp = INT8_MIN;
        }
        r.values[i] = (int8_t)tmp;
    }
    return r;
}

#endif  // VADD_H
