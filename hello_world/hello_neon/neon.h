// 2023-04-14 10:48
#ifndef COMMON_H
#define COMMON_H

#ifdef __aarch64__
#include <arm_neon.h>
#endif

#ifdef __x86_64
#include <neon_emu.h>
#endif

#ifdef __mips_msa
#include <neon_emu.h>
#endif

#endif  // COMMON_H
