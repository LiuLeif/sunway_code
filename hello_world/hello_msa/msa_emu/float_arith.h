// 2023-04-28 11:16
#ifndef FLOAT_ARITH_H
#define FLOAT_ARITH_H
#include <fenv.h>
#include <math.h>

#include "msa_emu_types.h"

v4f32 __msa_fadd_w(v4f32 a, v4f32 b) {
    v4f32 r;
    for (int i = 0; i < 4; i++) {
        r.values[i] = a.values[i] + b.values[i];
    }
    return r;
}

v4f32 __msa_fmadd_w(v4f32 c, v4f32 a, v4f32 b) {
    v4f32 r;
    for (int i = 0; i < 4; i++) {
        r.values[i] = (double)c.values[i] + (double)a.values[i] * b.values[i];
    }
    return r;
}

v4f32 __msa_frint_w(v4f32 a) {
    v4f32 r;
    fesetround(FE_TOWARDZERO);
    for (int i = 0; i < 4; i++) {
        float tmp = a.values[i];
        if (tmp > 0.0f) {
            tmp += 0.5f;
        } else {
            tmp -= 0.5f;
        }
        r.values[i] = nearbyintf(tmp);
    }
    fesetround(FE_TONEAREST);
    return r;
}

#endif  // FLOAT_ARITH_H
