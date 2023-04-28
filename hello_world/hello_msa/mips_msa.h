// 2023-04-14 10:48
#ifndef COMMON_H
#define COMMON_H

#ifndef EMU
#include <msa.h>
#else
#include <msa_emu.h>
#endif
#include <stdint.h>

typedef union {
    v16i8 i8;
    v16u8 u8;
    v4i32 i32;
    v4f32 f32;
} m128;

#endif  // COMMON_H
