// 2023-04-18 18:41
#include <neon.h>
#include <neon_test.h>
// int8x8_t vshl_s8(int8x8_t a,int8x8_t b)
// int16x4_t vshl_s16(int16x4_t a,int16x4_t b)
// int32x2_t vshl_s32(int32x2_t a,int32x2_t b)
// int64x1_t vshl_s64(int64x1_t a,int64x1_t b)
// uint8x8_t vshl_u8(uint8x8_t a,int8x8_t b)
// uint16x4_t vshl_u16(uint16x4_t a,int16x4_t b)
// uint32x2_t vshl_u32(uint32x2_t a,int32x2_t b)
// uint64x1_t vshl_u64(uint64x1_t a,int64x1_t b)
//
// int8x16_t vshlq_s8(int8x16_t a,int8x16_t b)
// int16x8_t vshlq_s16(int16x8_t a,int16x8_t b)
// int32x4_t vshlq_s32(int32x4_t a,int32x4_t b)
// int64x2_t vshlq_s64(int64x2_t a,int64x2_t b)
// uint8x16_t vshlq_u8(uint8x16_t a,int8x16_t b)
// uint16x8_t vshlq_u16(uint16x8_t a,int16x8_t b)
// uint32x4_t vshlq_u32(uint32x4_t a,int32x4_t b)
// uint64x2_t vshlq_u64(uint64x2_t a,int64x2_t b)
// ----------------------------------------------
// int64_t vshld_s64(int64_t a,int64_t b)
// uint64_t vshld_u64(uint64_t a,int64_t b)
// ----------------------------------------------
// int8x8_t vshl_n_s8(int8x8_t a,const int n)
// int16x4_t vshl_n_s16(int16x4_t a,const int n)
// int32x2_t vshl_n_s32(int32x2_t a,const int n)
// int64x1_t vshl_n_s64(int64x1_t a,const int n)
// uint8x8_t vshl_n_u8(uint8x8_t a,const int n)
// uint16x4_t vshl_n_u16(uint16x4_t a,const int n)
// uint32x2_t vshl_n_u32(uint32x2_t a,const int n)
// uint64x1_t vshl_n_u64(uint64x1_t a,const int n)
//
// int8x16_t vshlq_n_s8(int8x16_t a,const int n)
// int16x8_t vshlq_n_s16(int16x8_t a,const int n)
// int32x4_t vshlq_n_s32(int32x4_t a,const int n)
// int64x2_t vshlq_n_s64(int64x2_t a,const int n)
// uint8x16_t vshlq_n_u8(uint8x16_t a,const int n)
// uint16x8_t vshlq_n_u16(uint16x8_t a,const int n)
// uint32x4_t vshlq_n_u32(uint32x4_t a,const int n)
// uint64x2_t vshlq_n_u64(uint64x2_t a,const int n)
// ------------------------------------------------
// int64_t vshld_n_s64(int64_t a,const int n)
// uint64_t vshld_n_u64(uint64_t a,const int n)
TEST_CASE(test_vshl_s16) {
    struct {
        int16_t a[4];
        int16_t b[4];
        int16_t r[4];
    } test_vec[] = {
        {{22332, -2389, -6176, 24298},
         {-11, -12, 535, 14},
         {10, -1, 0, INT16_MIN}},
        {{-30833, -3392, 7263, 1769},
         {15, -10006, -26926, -19506},
         {INT16_MIN, -1, 0, 0}},
        {{24162, 31020, -13216, -28236},
         {9, 11, 16, 2895},
         {-15360, 24576, 0, 0}},
        {{-844, -30766, -24430, 3898}, {9, -9, -16, 2356}, {26624, -61, -1, 0}},
        {{7438, -17736, 2143, -16698}, {-10, -11, -15, 12}, {7, -9, 0, 24576}},
        {{2818, -21698, 29388, 22965},
         {10, 27061, 13, -13277},
         {2048, -1, INT16_MIN, 0}},
        {{17345, -26964, 10446, -1539},
         {14, -7811, -9, 19084},
         {16384, 0, 20, -1}},
        {{22929, -270, 5830, 15563},
         {10, -15886, -13, 12},
         {17408, -1, 0, -20480}},
    };

    for (size_t i = 0; i < (sizeof(test_vec) / sizeof(test_vec[0])); i++) {
        int16x4_t a = vld1_s16(test_vec[i].a);
        int16x4_t b = vld1_s16(test_vec[i].b);
        int16x4_t r = vshl_s16(a, b);
        int16x4_t check = vld1_s16(test_vec[i].r);
        ASSERT_EQUAL(r, check);
    }
    return 0;
}

TEST_CASE(test_simde_vshl_u8) {
    struct {
        uint8_t a[8];
        int8_t b[8];
        uint8_t r[8];
    } test_vec[] = {
        {{175, 152, 126, 1, 164, 17, 164, 72},
         {-8, 7, -8, -7, -7, -8, 7, 6},
         {0, 0, 0, 0, 1, 0, 0, 0}},
        {{189, 130, 234, 197, 247, 15, 90, 166},
         {-6, -40, -89, 4, -6, 4, -7, -5},
         {2, 0, 0, 80, 3, 240, 0, 5}},
        {{173, 5, 173, 224, 34, 193, 253, 223},
         {68, -6, -6, 7, -5, -1, -30, -7},
         {0, 0, 2, 0, 1, 96, 0, 1}},
        {{149, 67, 249, 57, 39, 110, 16, 213},
         {7, -5, -75, -7, INT8_MAX, -5, 7, -61},
         {128, 2, 0, 0, 0, 3, 0, 0}},
        {{106, 26, 178, 63, 152, 78, 222, 45},
         {-7, -5, 6, -5, 70, 119, -8, -5},
         {0, 0, 128, 1, 0, 0, 0, 1}},
        {{225, 119, 36, 251, 88, 84, 236, 195},
         {6, -7, 2, 7, -5, -8, 52, 7},
         {64, 0, 144, 128, 2, 0, 0, 128}},
        {{11, 11, 253, 2, 210, 118, 148, 179},
         {-19, -5, -82, 8, 12, -7, 4, 7},
         {0, 0, 0, 0, 0, 0, 64, 128}},
        {{82, 221, 165, 101, 164, 95, 174, 175},
         {106, -6, -5, 60, 8, 8, -5, 7},
         {0, 3, 5, 0, 0, 0, 5, 128}},
    };

    for (size_t i = 0; i < (sizeof(test_vec) / sizeof(test_vec[0])); i++) {
        uint8x8_t a = vld1_u8(test_vec[i].a);
        int8x8_t b = vld1_s8(test_vec[i].b);
        uint8x8_t r = vshl_u8(a, b);
        uint8x8_t check = vld1_u8(test_vec[i].r);
        ASSERT_EQUAL(r, check);
    }
    return 0;
}

TEST_CASE(test_vshlq_s16) {
    struct {
        int16_t a[8];
        int16_t b[8];
        int16_t r[8];
    } test_vec[] = {
        {{20268, 24220, 20072, -27645, 1744, 22176, -26671, 22566},
         {11, -13, -14, -13, -14, 16, 10, 4205},
         {24576, 2, 1, -4, 0, 0, 17408, 0}},
        {{-136, 18814, -23402, 17825, -12846, 12767, 24470, 16845},
         {-12, 11, -13, 12, -15, 12, 8, 14},
         {-1, -4096, -3, 4096, -1, -4096, -27136, 16384}},
        {{-469, -16002, 19293, -31486, 26094, -23408, -21452, 3518},
         {-12, 12, -26906, 10, 12, -6652, 8, -5838},
         {-1, -8192, 0, 2048, -8192, 18688, 13312, 0}},
        {{-27631, 18062, 19520, 30035, -28177, -10842, 18727, -27649},
         {937, 12, -11, 8, -12, 22726, -13913, -9},
         {-1, -8192, 9, 21248, -7, -1, 0, -55}},
        {{-17949, 3042, -7934, -21346, 6373, 2835, 25796, 8275},
         {12, 10, -9, 16574, -17268, -21795, -16, -3402},
         {12288, -30720, -16, -1, 0, 0, 0, 0}},
        {{27571, 30570, -16945, 11928, 4566, -18125, -3636, 23033},
         {-11, -17405, -9, 14, -11, -13, 14, -12},
         {13, -17584, -34, 0, 2, -3, 0, 5}},
        {{4379, -6174, -9470, -20416, 17330, 31085, 7165, 20457},
         {13, 14, 8, 11, 31451, 11, -13, -13},
         {24576, INT16_MIN, 512, 0, 0, 26624, 0, 2}},
        {{6155, 2113, 10804, -1193, 29232, -21751, 4536, -27899},
         {15, 16, 6903, -9, 13, 27370, -14, 2354},
         {INT16_MIN, 0, 21, -3, 0, -1, 0, 0}},
    };

    for (size_t i = 0; i < (sizeof(test_vec) / sizeof(test_vec[0])); i++) {
        int16x8_t a = vld1q_s16(test_vec[i].a);
        int16x8_t b = vld1q_s16(test_vec[i].b);
        int16x8_t r = vshlq_s16(a, b);
        int16x8_t check = vld1q_s16(test_vec[i].r);
        ASSERT_EQUAL(r, check);
    }
    return 0;
}
