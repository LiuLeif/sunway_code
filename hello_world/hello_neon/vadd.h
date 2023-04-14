// 2023-04-14 12:20
#ifndef VADD_H
#define VADD_H
#include "neon_emu_common.h"
// NOTE: 使用 64-bit vector 计算 a+b
int8x8_t vadd_s8(int8x8_t a, int8x8_t b) {
    int8x8_t r;
    for (int i = 0; i < 8; i++) {
        r.values[i] = a.values[i] + b.values[i];
    }
    return r;
}
// NOTE: 使用 128-bit vector 计算 a+b (q 后缀表示 128-bit)
int8x16_t vaddq_s8(int8x16_t a, int8x16_t b) {
    int8x16_t r;
    for (int i = 0; i < 16; i++) {
        r.values[i] = a.values[i] + b.values[i];
    }
    return r;
}
#endif  // VADD_H
