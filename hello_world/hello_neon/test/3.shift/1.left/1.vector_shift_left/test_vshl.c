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
