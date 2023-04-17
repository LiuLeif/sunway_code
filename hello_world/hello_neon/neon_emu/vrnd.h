// 2023-04-17 18:36
#ifndef VRND_H
#define VRND_H

#include <math.h>
#include <stdint.h>

#include "neon_emu_types.h"

float32x2_t vrnd_f32(float32x2_t a) {
    float32x2_t r;
    for (int i = 0; i < 2; i++) {
        r.values[i] =
            a.values[i] > 0.0f ? floorf(a.values[i]) : ceilf(a.values[i]);
    }
    return r;
}

float32x2_t vrndm_f32(float32x2_t a) {
    float32x2_t r;
    for (int i = 0; i < 2; i++) {
        r.values[i] = floorf(a.values[i]);
    }
    return r;
}
#endif  // VRND_H
