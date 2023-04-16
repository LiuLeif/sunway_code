// 2023-04-14 13:07
#ifndef VLD_H
#define VLD_H
#include "neon_emu_types.h"
int8x8_t vld1_s8(const int8_t* ptr) {
    int8x8_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

uint8x8_t vld1_u8(const uint8_t* ptr) {
    uint8x8_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

int8x16_t vld1q_s8(const int8_t* ptr) {
    int8x16_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

int16x8_t vld1q_s16(const int16_t* ptr) {
    int16x8_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

float32x2_t vld1_f32(const float* ptr) {
    float32x2_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

float32x4_t vld1q_f32(const float* ptr) {
    float32x4_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

#endif  // VLD_H
