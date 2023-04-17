// 2023-04-14 10:48
#ifndef COMMON_H
#define COMMON_H
#include <assert.h>
#include <math.h>
#include <stddef.h>
#include <stdio.h>

#ifndef EMU
#include <arm_neon.h>
#define ASSERT_EQUAL(n, a, b)         \
    do {                              \
        for (int i = 0; i < n; i++) { \
            assert(a[i] == b[i]);     \
        }                             \
    } while (0)

#define ASSERT_CLOSE(n, a, b)                         \
    do {                                              \
        for (int i = 0; i < n; i++) {                 \
            if (isnanf(a[i]) || isnanf(b[i])) {       \
                assert(isnanf(b[i]) && isnanf(a[i])); \
            } else {                                  \
                assert(fabs(a[i] - b[i]) < 1e-2);     \
            }                                         \
        }                                             \
    } while (0)
#else
#include "./neon_emu.h"
#define ASSERT_EQUAL(n, a, b)                   \
    do {                                        \
        for (int i = 0; i < n; i++) {           \
            assert(a.values[i] == b.values[i]); \
        }                                       \
    } while (0)

#define ASSERT_CLOSE(n, a, b)                                       \
    do {                                                            \
        for (int i = 0; i < n; i++) {                               \
            if (isnanf(a.values[i]) || isnanf(b.values[i])) {       \
                assert(isnanf(b.values[i]) && isnanf(a.values[i])); \
            } else {                                                \
                assert(fabs(a.values[i] - b.values[i]) < 1e-2);     \
            }                                                       \
        }                                                           \
    } while (0)
#endif

#define ASSERT_EQUAL_SCALAR(a, b) \
    do {                          \
        assert(a == b);           \
    } while (0)

#define ASSERT_CLOSE_SCALAR(a, b)   \
    do {                            \
        assert(fabs(a - b) < 1e-2); \
    } while (0)

#define TEST_CASE(name) __attribute__((constructor)) int name()

int main(int argc, char *argv[]) { return 0; }
#endif  // COMMON_H
