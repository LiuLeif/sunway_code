// 2023-04-14 13:08
#ifndef ARM_NEON_COMMON_H
#define ARM_NEON_COMMON_H

#include <stdint.h>
#include <string.h>

typedef union {
    int8_t values[8];
} int8x8_t;

typedef union {
    uint8_t values[8];
} uint8x8_t;

typedef union {
    int8_t values[16];
} int8x16_t;

typedef union {
    int16_t values[8];
} int16x8_t;

typedef union {
    float values[2];
} float32x2_t;

typedef union {
    float values[4];
} float32x4_t;

typedef union {
    int16_t values[4];
} int16x4_t;

typedef union {
    int32_t values[4];
} int32x4_t;

#endif  // ARM_NEON_COMMON_H
