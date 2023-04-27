// 2023-04-14 13:08
#ifndef ARM_NEON_COMMON_H
#define ARM_NEON_COMMON_H

#include <stdint.h>
#include <string.h>

#define DEF_TYPE(base, n, abbr) \
    typedef union {             \
        base##_t values[n];     \
    } v##n##abbr;

DEF_TYPE(int8, 16, i8);
DEF_TYPE(uint8, 16, u8);
DEF_TYPE(int16, 8, i16);

#endif  // ARM_NEON_COMMON_H
