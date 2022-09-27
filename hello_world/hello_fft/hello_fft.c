// 2022-09-27 15:24
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "kissfft/kiss_fft.h"

#define PI 3.14159
#define N 128

void my_fft(int n_point, kiss_fft_cpx *in, kiss_fft_cpx *out) {
    for (int i = 0; i < n_point; i++) {
        float tmp_r = 0.0f;
        float tmp_i = 0.0f;
        for (int k = 0; k < n_point; k++) {
            tmp_r += in[k].r * cos(i * 2.0 * PI * k / n_point);
            tmp_i -= in[k].r * sin(i * 2.0 * PI * k / n_point);
        }
        out[i].r = tmp_r;
        out[i].i = tmp_i;
    }
}

void my_ifft(int n_point, kiss_fft_cpx *in, kiss_fft_cpx *out) {
    for (int i = 0; i < n_point; i++) {
        float tmp = 0.0f;
        for (int k = 0; k < n_point; k++) {
            tmp += in[k].r * cos(2.0 * PI * i * k / n_point);
            tmp -= in[k].i * sin(2.0 * PI * i * k / n_point);
        }
        out[i].r = tmp;
        out[i].i = 0.0f;
    }
}

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
    printf("fft:\n");
    kiss_fft(cfg, in, out);
    printf("1hz cos: %f\n", out[1].r / 64);
    printf("2hz sin: %f\n", -out[2].i / 64);
    printf("------\n");

    printf("fft manually:\n");
    my_fft(N, in, out);
    printf("1hz cos: %f\n", out[1].r / 64);
    printf("2hz sin: %f\n", -out[2].i / 64);
    printf("------\n");

    printf("ifft:\n");
    cfg = kiss_fft_alloc(N, 1, 0, 0);
    kiss_fft(cfg, out, in_restored);
    for (int i = 0; i < 10; i++) {
        printf("%f %f\n", in_restored[i].r / N, in[i].r);
    }
    printf("------\n");

    printf("ifft manually:\n");
    my_ifft(N, out, in_restored);
    for (int i = 0; i < 10; i++) {
        printf("%f %f\n", in_restored[i].r / N, in[i].r);
    }
}
