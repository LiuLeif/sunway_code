// 2023-04-14 12:20
#ifndef VADD_H
#define VADD_H

#include <neon_emu_types.h>

#include "util.h"

int8x8_t vadd_s8(int8x8_t a, int8x8_t b) {
    int8x8_t r;
    r.v.i8 = __msa_addv_b(a.v.i8, b.v.i8);
    return r;
}

int8x16_t vaddq_s8(int8x16_t a, int8x16_t b) {
    int8x16_t r;
    r.v.i8 = __msa_addv_b(a.v.i8, b.v.i8);
    return r;
}

int8x8_t vaddhn_s16(int16x8_t a, int16x8_t b) {
    int16x8_t tmp;
    tmp.v.i16 = __msa_addv_h(a.v.i16, b.v.i16);
    tmp.v.i16 = __msa_srai_h(tmp.v.i16, 8);

    int8x8_t r;
    COPY(r, tmp);
    return r;
}

int8x8_t vhadd_s8(int8x8_t a, int8x8_t b) {
    int16x8_t _a, _b, tmp;
    COPY(_a, a);
    COPY(_b, b);

    tmp.v.i16 = __msa_addv_h(_a.v.i16, _b.v.i16);
    tmp.v.i16 = __msa_srai_h(tmp.v.i16, 1);

    int8x8_t r;
    COPY(r, tmp);

    return r;
}

int8x8_t vrhadd_s8(int8x8_t a, int8x8_t b) {
    int16x8_t _a, _b, tmp;
    COPY(_a, a);
    COPY(_b, b);

    tmp.v.i16 = __msa_addv_h(_a.v.i16, _b.v.i16);
    tmp.v.i16 = __msa_addvi_h(tmp.v.i16, 1);
    tmp.v.i16 = __msa_srai_h(tmp.v.i16, 1);

    int8x8_t r;
    COPY(r, tmp);
    return r;
}

int8x8_t vqadd_s8(int8x8_t a, int8x8_t b) {
    int8x8_t r;
    r.v.i8 = __msa_adds_s_b(a.v.i8, b.v.i8);
    return r;
}

int8x8_t vuqadd_s8(int8x8_t a, uint8x8_t b) {
    int8x8_t r;
    int16x8_t _a, _b, _r;
    COPY(_a, a);
    COPY(_b, b);
    _r.v.i16 = __msa_addv_h(_a.v.i16, _b.v.i16);
    _r.v.i16 = __msa_sat_s_h(_r.v.i16, 7);
    COPY(r, _r);
    return r;
}

uint8x8_t vsqadd_u8(uint8x8_t a, int8x8_t b) {
    uint8x8_t r;
    int16x8_t _a, _b, _r;
    COPY(_a, a);
    COPY(_b, b);
    _r.v.i16 = __msa_addv_h(_a.v.i16, _b.v.i16);
    for (int i = 0; i < 8; i++) {
        if (_r.values[i] > UINT8_MAX) {
            _r.values[i] = UINT8_MAX;
        }
        if (_r.values[i] < 0) {
            _r.values[i] = 0;
        }
    }
    COPY(r, _r);
    return r;
}

int8_t vqaddb_s8(int8_t a, int8_t b) {
    int16_t r = (int16_t)a + b;
    if (r > INT8_MAX) {
        r = INT8_MAX;
    }
    if (r < INT8_MIN) {
        r = INT8_MIN;
    }
    return (int8_t)r;
}

int8_t vuqaddb_s8(int8_t a, uint8_t b) {
    int16_t r = (int16_t)a + b;
    if (r > INT8_MAX) {
        r = INT8_MAX;
    }
    if (r < INT8_MIN) {
        r = INT8_MIN;
    }
    return (int8_t)r;
}

uint8_t vsqaddb_u8(uint8_t a, int8_t b) {
    int16_t r = (int16_t)a + b;
    if (r > UINT8_MAX) {
        r = UINT8_MAX;
    }
    if (r < 0) {
        r = 0;
    }
    return (uint8_t)r;
}

int16x8_t vaddl_s8(int8x8_t a, int8x8_t b) {
    int16x8_t r;
    int16x8_t _a, _b;
    COPY(_a, a);
    COPY(_b, b);
    r.v.i16 = __msa_addv_h(_a.v.i16, _b.v.i16);
    return r;
}

int16x8_t vaddl_high_s8(int8x16_t a, int8x16_t b) {
    int16x8_t r;
    int16x8_t _a, _b;
    COPY_HIGH(_a, a);
    COPY_HIGH(_b, b);
    r.v.i16 = __msa_addv_h(_a.v.i16, _b.v.i16);
    return r;
}

int16x8_t vaddw_s8(int16x8_t a, int8x8_t b) {
    int16x8_t r;
    int16x8_t _b;
    COPY(_b, b);
    r.v.i16 = __msa_addv_h(a.v.i16, _b.v.i16);
    return r;
}

int16x8_t vaddw_high_s8(int16x8_t a, int8x16_t b) {
    int16x8_t r;
    int16x8_t _b;
    COPY_HIGH(_b, b);
    r.v.i16 = __msa_addv_h(a.v.i16, _b.v.i16);
    return r;
}

#endif  // VADD_H
