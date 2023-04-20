// 2023-04-20 11:15
#include <neon.h>
#include <neon_test.h>
// int8x8_t vrshrn_n_s16(int16x8_t a,const int n)
// int16x4_t vrshrn_n_s32(int32x4_t a,const int n)
// int32x2_t vrshrn_n_s64(int64x2_t a,const int n)
// uint8x8_t vrshrn_n_u16(uint16x8_t a,const int n)
// uint16x4_t vrshrn_n_u32(uint32x4_t a,const int n)
// uint32x2_t vrshrn_n_u64(uint64x2_t a,const int n)
// ---------------------------------------------------------------
// int8x16_t vrshrn_high_n_s16(int8x8_t r,int16x8_t a,const int n)
// int16x8_t vrshrn_high_n_s32(int16x4_t r,int32x4_t a,const int n)
// int32x4_t vrshrn_high_n_s64(int32x2_t r,int64x2_t a,const int n)
// uint8x16_t vrshrn_high_n_u16(uint8x8_t r,uint16x8_t a,const int n)
// uint16x8_t vrshrn_high_n_u32(uint16x4_t r,uint32x4_t a,const int n)
// uint32x4_t vrshrn_high_n_u64(uint32x2_t r,uint64x2_t a,const int n)

TEST_CASE(test_vrshrn_n_s16) {
    static const struct {
        int16_t a[8];
        int8_t r1[16];
        int8_t r3[16];
        int8_t r5[16];
        int8_t r6[16];
        int8_t r8[16];
    } test_vec[] = {
        {{10562, 6844, 14407, -31970, 8622, -13971, 17167, -26709},
         {-95, 94, 36, -113, -41, -73, -120, -42},
         {40, 88, 9, 100, 54, 46, 98, -11},
         {74, -42, -62, 25, 13, 75, 24, -67},
         {-91, 107, -31, 12, -121, 38, 12, 95},
         {41, 27, 56, -125, 34, -55, 67, -104}},
        {{30334, -3256, 2880, -20934, 21822, 14204, -24458, -18006},
         {63, -92, -96, 29, -97, -66, 59, -43},
         {-48, 105, 104, -57, -88, -16, 15, 53},
         {-76, -102, 90, 114, -86, -68, 4, -51},
         {-38, -51, 45, -71, 85, -34, -126, -25},
         {118, -13, 11, -82, 85, 55, -96, -70}},
        {{26314, 4563, -3682, 19604, 275, 8726, -16060, -15687},
         {101, -22, -49, 74, -118, 11, -94, 93},
         {-39, 58, 52, -109, 34, 67, 41, 87},
         {54, -113, -115, 101, 9, 17, 10, 22},
         {-101, 71, -58, 50, 4, -120, 5, 11},
         {103, 18, -14, 77, 1, 34, -63, -61}},
        {{311, 30645, -4339, 19237, -24252, -17533, 11586, 3188},
         {-100, -37, -121, -109, -94, -62, -95, 58},
         {39, -9, -30, 101, 41, 112, -88, -113},
         {10, -66, 120, 89, 10, -36, 106, 100},
         {5, -33, -68, 45, -123, -18, -75, 50},
         {1, 120, -17, 75, -95, -68, 45, 12}},
        {{18324, 12829, -20168, 19327, -27214, -2451, 10070, -29256},
         {-54, 15, -100, -64, -39, 55, -85, -36},
         {-13, 68, 39, 112, -74, -50, -21, -73},
         {61, -111, -118, 92, -82, -77, 59, 110},
         {30, -56, -59, 46, 87, -38, -99, 55},
         {72, 50, -79, 75, -106, -10, 39, -114}},
        {{27944, 13572, 10588, -24191, 1227, 3420, -12239, -15079},
         {-108, -126, -82, -63, 102, -82, 25, -115},
         {-91, -95, 44, 48, -103, -84, 6, -93},
         {105, -88, 75, 12, 38, 107, -126, 41},
         {-75, -44, -91, -122, 19, 53, 65, 20},
         {109, 53, 41, -94, 5, 13, -48, -59}},
        {{13847, 20472, 30695, -26213, 2060, 25231, 18223, 22767},
         {12, -4, -12, -50, 6, 72, -104, 120},
         {-61, -1, -3, 51, 2, 82, -26, 30},
         {-79, INT8_MIN, -65, -51, 64, 20, 57, -57},
         {-40, 64, -32, 102, 32, -118, 29, 100},
         {54, 80, 120, -102, 8, 99, 71, 89}},
        {{-3148, 4237, 3612, -6223, 3346, 17652, 3549, -3063},
         {-38, 71, 14, -39, -119, 122, -17, 5},
         {119, 18, -60, -10, -94, -97, -68, -127},
         {-98, -124, 113, 62, 105, 40, 111, -96},
         {-49, 66, 56, -97, 52, 20, 55, -48},
         {-12, 17, 14, -24, 13, 69, 14, -12}},
    };

    for (size_t i = 0; i < (sizeof(test_vec) / sizeof(test_vec[0])); i++) {
        int16x8_t a = vld1q_s16(test_vec[i].a);

        int8x8_t r1 = vrshrn_n_s16(a, 1);
        int8x8_t r3 = vrshrn_n_s16(a, 3);
        int8x8_t r5 = vrshrn_n_s16(a, 5);
        int8x8_t r6 = vrshrn_n_s16(a, 6);
        int8x8_t r8 = vrshrn_n_s16(a, 8);
        
        int8x8_t check1 = vld1_s8(test_vec[i].r1);
        int8x8_t check3 = vld1_s8(test_vec[i].r3);
        int8x8_t check5 = vld1_s8(test_vec[i].r5);
        int8x8_t check6 = vld1_s8(test_vec[i].r6);
        int8x8_t check8 = vld1_s8(test_vec[i].r8);

        ASSERT_EQUAL(r1, check1);
        ASSERT_EQUAL(r3, check3);
        ASSERT_EQUAL(r5, check5);
        ASSERT_EQUAL(r6, check6);
        ASSERT_EQUAL(r8, check8);
    }

    return 0;
}
