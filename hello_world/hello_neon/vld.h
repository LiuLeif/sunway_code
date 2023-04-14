// 2023-04-14 13:07
#ifndef VLD_H
#define VLD_H
#include "neon_emu_common.h"
int8x8_t vld1_s8(const int8_t* ptr) {
    int8x8_t r;
    memcpy(&r, ptr, 8);
    return r;
}

int8x16_t vld1q_s8(const int8_t* ptr) {
    int8x16_t r;
    memcpy(&r, ptr, 16);
    return r;
}
#endif  // VLD_H
