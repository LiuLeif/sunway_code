// 2023-04-17 15:08
#ifndef REINTERP_H
#define REINTERP_H
#include <stdint.h>

#include "neon_emu_types.h"

poly8x8_t vreinterpret_p8_u8(uint8x8_t a) {
    poly8x8_t r;
    memcpy(&r, &a, sizeof(a));
    return r;
}

#endif  // REINTERP_H
