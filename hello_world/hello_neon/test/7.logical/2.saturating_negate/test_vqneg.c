// 2023-04-20 16:48
#include <neon.h>
#include <neon_test.h>
// int8x8_t vqneg_s8(int8x8_t a)
// int16x4_t vqneg_s16(int16x4_t a)
// int32x2_t vqneg_s32(int32x2_t a)
// int64x1_t vqneg_s64(int64x1_t a)
//
// int8x16_t vqnegq_s8(int8x16_t a)
// int16x8_t vqnegq_s16(int16x8_t a)
// int32x4_t vqnegq_s32(int32x4_t a)
// int64x2_t vqnegq_s64(int64x2_t a)
// ---------------------------------
// int8_t vqnegb_s8(int8_t a)
// int16_t vqnegh_s16(int16_t a)
// int32_t vqnegs_s32(int32_t a)
// int64_t vqnegd_s64(int64_t a)

TEST_CASE(test_vqnegb_s8) {
    static const struct {
        int8_t a;
        int8_t r;
    } test_vec[] = {
        {INT8_MIN, INT8_MAX},
        {-59, 59},
        {-53, 53},
        {96, -96},
        {75, -75},
        {-55, 55},
        {-47, 47},
        {-61, 61}};

    for (size_t i = 0; i < (sizeof(test_vec) / sizeof(test_vec[0])); i++) {
        int8_t r = vqnegb_s8(test_vec[i].a);
        ASSERT_EQUAL_SCALAR(r, test_vec[i].r);
    }
    return 0;
}
