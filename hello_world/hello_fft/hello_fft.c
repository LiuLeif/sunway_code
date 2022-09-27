// 2022-09-27 15:24
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "kissfft/kiss_fft.h"

#define PI 3.14159

int main(int argc, char *argv[]) {
    kiss_fft_cfg cfg = kiss_fft_alloc(128, 0, 0, 0);
    kiss_fft_cpx *in =
        (kiss_fft_cpx *)KISS_FFT_MALLOC(128 * sizeof(kiss_fft_cpx));
    kiss_fft_cpx *out =
        (kiss_fft_cpx *)KISS_FFT_MALLOC(128 * sizeof(kiss_fft_cpx));
    /* NOTE: 1hz cos with amplitude 1 + 2hz sin with amplitude 10 */
    for (int i = 0; i < 128; i++) {
        in[i].r =
            cos(i * 2.0 * PI / 128.0) + 10 * sin(2 * i * 2.0 * PI / 128.0);
    }
    kiss_fft(cfg, in, out);
    printf("1hz cos: %f\n", out[1].r / 64);
    printf("2hz sin: %f\n", -out[2].i / 64);

    float a = 0.0;
    for (int i = 0; i < 128; i++) {
        a += in[i].r * cos(2 * PI * i / 128.0);
    }
    printf("1hz cos: %f\n", a / 64);
    float b = 0.0;
    for (int i = 0; i < 128; i++) {
        b += in[i].r * sin(2.0 * PI * 2.0 * i / 128.0);
    }
    printf("2hz sin: %f\n", b / 64);
}
