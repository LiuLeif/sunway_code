// 2023-04-20 13:58
#ifndef VCVT_H
#define VCVT_H
#include <fenv.h>
#include <math.h>

#include <neon_emu_types.h>

int32x2_t vcvt_s32_f32(float32x2_t a) {
    int32x2_t r;
    for (int i = 0; i < 2; i++) {
        r.values[i] = a.values[i];
    }
    return r;
}

int32x2_t vcvta_s32_f32(float32x2_t a) {
    int32x2_t r;
    fesetround(FE_TOWARDZERO);
    float32_t tmp;
    for (int i = 0; i < 2; i++) {
        tmp = a.values[i];
        if (tmp >= 0.0f) {
            tmp += 0.5f;
        } else {
            tmp -= 0.5f;
        }
        r.values[i] = nearbyintf(tmp);
    }
    return r;
}

#endif  // VCVT_H
