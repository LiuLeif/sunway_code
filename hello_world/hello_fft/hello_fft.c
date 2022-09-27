// 2022-09-27 15:24
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "kissfft/kiss_fft.h"

#define PI 3.14159
#define N 128

int main(int argc, char *argv[]) {
    kiss_fft_cfg cfg = kiss_fft_alloc(N, 0, 0, 0);
    kiss_fft_cpx *in =
        (kiss_fft_cpx *)KISS_FFT_MALLOC(N * sizeof(kiss_fft_cpx));
    kiss_fft_cpx *out =
        (kiss_fft_cpx *)KISS_FFT_MALLOC(N * sizeof(kiss_fft_cpx));
    kiss_fft_cpx *in_restored =
        (kiss_fft_cpx *)KISS_FFT_MALLOC(N * sizeof(kiss_fft_cpx));
    /* NOTE: 1hz cos with amplitude 1 + 2hz sin with amplitude 10 */
    for (int i = 0; i < N; i++) {
        in[i].r = cos(i * 2.0 * PI / N) + 10 * sin(2 * i * 2.0 * PI / N);
    }
    printf("------\n");
    printf("using kissfft:\n");
    kiss_fft(cfg, in, out);
    printf("1hz cos: %f\n", out[1].r / 64);
    printf("2hz sin: %f\n", -out[2].i / 64);

    printf("------\n");
    printf("calc manually:\n");
    float a = 0.0;
    for (int i = 0; i < N; i++) {
        a += in[i].r * cos(2 * PI * i / N);
    }
    printf("1hz cos: %f\n", a / 64);
    float b = 0.0;
    for (int i = 0; i < N; i++) {
        b += in[i].r * sin(2.0 * PI * 2.0 * i / N);
    }
    printf("2hz sin: %f\n", b / 64);
    printf("------\n");
    printf("ifft:\n");
    cfg = kiss_fft_alloc(N, 1, 0, 0);
    kiss_fft(cfg, out, in_restored);
    for (int i = 0; i < 10; i++) {
        printf("%f %f\n", in_restored[i].r / N, in[i].r);
    }
    printf("------\n");
    printf("ifft manually:\n");
    for (int i = 0; i < 10; i++) {
        float x_restored = 0;
        for (int k = 0; k < N; k++) {
            x_restored += out[k].r * cos(2.0 * PI * i * k / N);
            x_restored -= out[k].i * sin(2.0 * PI * i * k / N);
        }
        printf("%f %f\n", x_restored / N, in[i].r);
    }
}
