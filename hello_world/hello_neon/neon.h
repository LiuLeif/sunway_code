// 2023-04-14 10:48
#ifndef COMMON_H
#define COMMON_H
#include <assert.h>
#include <stddef.h>
#include <stdio.h>

#ifndef EMU
#include <arm_neon.h>
#define ASSERT_EQUAL(n, a, b)     \
    for (int i = 0; i < n; i++) { \
        assert(a[i] == b[i]);     \
    }
#else
#include "./neon_emu.h"
#define ASSERT_EQUAL(n, a, b)               \
    for (int i = 0; i < n; i++) {           \
        assert(a.values[i] == b.values[i]); \
    }
#endif

#define TEST_CASE(name) __attribute__((constructor)) int name ()

int main(int argc, char *argv[]) { return 0; }
#endif  // COMMON_H
