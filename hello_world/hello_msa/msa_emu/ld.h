// 2023-04-26 16:46
#ifndef LD_H
#define LD_H

#include "msa_emu_types.h"

v16i8 __msa_ld_b(void *a, int n) {
    v16i8 r;
    memcpy(&r, a, 16);
    return r;
}
#endif  // LD_H
