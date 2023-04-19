// 2023-04-14 13:07
#ifndef VLD_H
#define VLD_H
#include "neon_emu_types.h"

int8x8_t vld1_s8(const int8_t* ptr) {
    int8x8_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

int16x4_t vld1_s16(const int16_t* ptr) {
    int16x4_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

uint8x8_t vld1_u8(const uint8_t* ptr) {
    uint8x8_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

uint16x4_t vld1_u16(const uint16_t* ptr) {
    uint16x4_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

uint32x2_t vld1_u32(const uint32_t* ptr) {
    uint32x2_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

uint64x1_t vld1_u64(const uint64_t* ptr) {
    uint64x1_t r;
    memcpy(&r, ptr, sizeof(r));
    return r;
}

int64x1_t vld1_s64(const int64_t* ptr) {
    int64x1_t r;
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

int32x4_t vld1q_s32(const int32_t* ptr) {
    int32x4_t r;
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
