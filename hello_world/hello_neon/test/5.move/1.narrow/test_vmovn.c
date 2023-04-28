// 2023-04-20 14:49
#include <neon.h>
#include <neon_test.h>
// int8x8_t vmovn_s16(int16x8_t a)
//              ^---narrow
// int16x4_t vmovn_s32(int32x4_t a)
// int32x2_t vmovn_s64(int64x2_t a)
// uint8x8_t vmovn_u16(uint16x8_t a)
// uint16x4_t vmovn_u32(uint32x4_t a)
// uint32x2_t vmovn_u64(uint64x2_t a)
//
// int8x16_t vmovn_high_s16(int8x8_t r,int16x8_t a)
// int16x8_t vmovn_high_s32(int16x4_t r,int32x4_t a)
// int32x4_t vmovn_high_s64(int32x2_t r,int64x2_t a)
// uint8x16_t vmovn_high_u16(uint8x8_t r,uint16x8_t a)
// uint16x8_t vmovn_high_u32(uint16x4_t r,uint32x4_t a)
// uint32x4_t vmovn_high_u64(uint32x2_t r,uint64x2_t a)

TEST_CASE(test_simde_vmovn_s16) {
    static const struct {
        int16_t a[8];
        int8_t r[8];
    } test_vec[] = {
        {{-9747, 2915, 12465, -19440, -27230, 26615, -10739, -1162},
         {-19, 99, -79, 16, -94, -9, 13, 118}},
        {{-4564, 11796, 560, -5089, -4592, -22646, 6419, 136},
         {44, 20, 48, 31, 16, -118, 19, -120}},
        {{-5134, -23541, 7196, -16808, 20657, -16602, -25562, 21178},
         {-14, 11, 28, 88, -79, 38, 38, -70}},
        {{-12661, -17536, -24623, -7769, 12941, -24440, 4171, 16032},
         {-117, INT8_MIN, -47, -89, -115, -120, 75, -96}},
        {{-21253, 6114, 15048, 31190, -886, -20424, -3432, 8962},
         {-5, -30, -56, -42, -118, 56, -104, 2}},
        {{-31807, -27937, -31198, -20365, -1096, 1104, -3829, 1602},
         {-63, -33, 34, 115, -72, 80, 11, 66}},
        {{9373, 25886, -2978, -5666, 6128, -30567, -25591, -13652},
         {-99, 30, 94, -34, -16, -103, 9, -84}},
        {{-29921, 16732, -12527, -13583, 17098, -10802, 4147, -12068},
         {31, 92, 17, -15, -54, -50, 51, -36}}};

    for (size_t i = 0; i < (sizeof(test_vec) / sizeof(test_vec[0])); i++) {
        int16x8_t a = vld1q_s16(test_vec[i].a);
        int8x8_t r = vmovn_s16(a);
        int8x8_t check = vld1_s8(test_vec[i].r);
        ASSERT_EQUAL(r, check);
    }
    return 0;
}

TEST_CASE(test_vmovn_high_s16) {
    static const struct {
        int8_t r[8];
        int16_t a[8];
        int8_t res[16];
    } test_vec[] = {
        {{-75, -23, -101, -44, -30, -4, -58, 42},
         {-2815, 20026, -27075, -9535, -19091, -13510, -18313, 21770},
         {-75, -23, -101, -44, -30, -4, -58, 42, 1, 58, 61, -63, 109, 58, 119,
          10}},
        {{77, -89, -42, 17, -17, 74, -118, -92},
         {9523, 5496, 16161, 8767, 31028, 29297, 12815, 31820},
         {77, -89, -42, 17, -17, 74, -118, -92, 51, 120, 33, 63, 52, 113, 15,
          76}},
        {{-25, -121, 71, 95, 63, 81, -76, -115},
         {-29959, -5986, 10452, 1933, 1357, 28188, 23364, 31120},
         {-25, -121, 71, 95, 63, 81, -76, -115, -7, -98, -44, -115, 77, 28, 68,
          -112}},
        {{-44, 1, -21, -29, 51, 55, 96, 27},
         {-22594, -390, 12025, -3445, 10681, -29222, 26449, -24939},
         {-44, 1, -21, -29, 51, 55, 96, 27, -66, 122, -7, -117, -71, -38, 81,
          -107}},
        {{109, -79, 12, -79, 13, -100, 42, -31},
         {5534, -11835, 9549, 3052, 26316, -15095, -27499, 20151},
         {109, -79, 12, -79, 13, -100, 42, -31, -98, -59, 77, -20, -52, 9, -107,
          -73}},
        {{-67, -110, -37, 14, -7, 112, -84, 102},
         {-18398, 12056, 16981, -3312, -10920, -23100, -19974, -14416},
         {-67, -110, -37, 14, -7, 112, -84, 102, 34, 24, 85, 16, 88, -60, -6,
          -80}},
        {{23, -70, -116, -84, 78, 68, -6, 12},
         {-10538, -12518, -14522, 26678, 20095, -11113, -22384, -5945},
         {23, -70, -116, -84, 78, 68, -6, 12, -42, 26, 70, 54, INT8_MAX, -105,
          -112, -57}},
        {{125, -116, -115, 120, 61, 62, 63, 84},
         {-13320, 17921, -1265, -6830, 28113, 6325, -5324, -19584},
         {125, -116, -115, 120, 61, 62, 63, 84, -8, 1, 15, 82, -47, -75, 52,
          INT8_MIN}}};

    for (size_t i = 0; i < (sizeof(test_vec) / sizeof(test_vec[0])); i++) {
        int8x8_t r = vld1_s8(test_vec[i].r);
        int16x8_t a = vld1q_s16(test_vec[i].a);
        int8x16_t res = vmovn_high_s16(r, a);
        int8x16_t check = vld1q_s8(test_vec[i].res);
        ASSERT_EQUAL(res, check);
    }
    return 0;
}