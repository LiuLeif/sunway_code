// 2023-04-26 14:41
#ifndef ADD_H
#define ADD_H

#include "msa_emu_types.h"

v16i8 __msa_addv_b(v16i8 a, v16i8 b) {
    v16i8 r;
    for (int i = 0; i < 16; i++) {
        r.values[i] = a.values[i] + b.values[i];
    }
    return r;
}

#endif  // ADD_H
