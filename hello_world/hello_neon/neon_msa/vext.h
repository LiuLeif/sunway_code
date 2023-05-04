// 2023-04-21 15:22
#ifndef VEXT_H
#define VEXT_H

#include <neon_emu_types.h>

int8x8_t vext_s8(int8x8_t a, int8x8_t b, int n) {
    int8x8_t r;
    r.v.i8 = __msa_sld_b(b.v.i8, a.v.i8, n);
    return r;
}
#endif  // VEXT_H
