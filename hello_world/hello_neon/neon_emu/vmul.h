// 2023-04-14 19:39
#ifndef VMUL_H
#define VMUL_H
#include <stdint.h>

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

float32x2_t vfma_f32(float32x2_t a, float32x2_t b, float32x2_t c) {
    float32x2_t r;
    for (int i = 0; i < 2; i++) {
        r.values[i] = (double)a.values[i] + (double)b.values[i] * c.values[i];
    }
    return r;
}

float32x2_t vfma_lane_f32(
    float32x2_t a, float32x2_t b, float32x2_t v, int lane) {
    float32x2_t r;
    for (int i = 0; i < 2; i++) {
        r.values[i] =
            (double)a.values[i] + (double)b.values[i] * v.values[lane];
    }
    return r;
}

float32x2_t vfma_laneq_f32(
    float32x2_t a, float32x2_t b, float32x4_t v, int lane) {
    float32x2_t r;
    for (int i = 0; i < 2; i++) {
        r.values[i] =
            (double)a.values[i] + (double)b.values[i] * v.values[lane];
    }
    return r;
}

float vfmas_lane_f32(float a, float b, float32x2_t v, int lane) {
    float r = (double)a + (double)b * v.values[lane];
    return r;
}

int16x4_t vqdmulh_s16(int16x4_t a, int16x4_t b) {
    int16x4_t r;
    for (int i = 0; i < 4; i++) {
        r.values[i] = (a.values[i] * b.values[i] * 2) >> 16;
    }
    return r;
}

int16x4_t vqrdmulh_s16(int16x4_t a, int16x4_t b) {
    int16x4_t r;
    for (int i = 0; i < 4; i++) {
        r.values[i] = (a.values[i] * b.values[i] * 2 + (1 << 15)) >> 16;
    }
    return r;
}

int32x4_t vqdmull_s16(int16x4_t a, int16x4_t b) {
    int32x4_t r;
    for (int i = 0; i < 4; i++) {
        r.values[i] = (a.values[i] * b.values[i] * 2);
    }
    return r;
}
#endif  // VMUL_H
