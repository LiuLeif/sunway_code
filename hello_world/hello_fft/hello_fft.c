// 2022-09-27 15:24
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

#include "kissfft/kiss_fft.h"

#define PI 3.14159
#define N 128

#define TIMEIT(F, REPS)                  \
    {                                    \
        clock_t start = clock();         \
        for (int i = 0; i < REPS; ++i) { \
            F;                           \
        }                                \
        clock_t diff = clock() - start;  \
        printf("%s: %ld\n", #F, diff);   \
    }

void my_dft(int n_point, kiss_fft_cpx *in, kiss_fft_cpx *out) {
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

void my_idft(int n_point, kiss_fft_cpx *in, kiss_fft_cpx *out) {
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

void my_fft(int n_point, kiss_fft_cpx *in, kiss_fft_cpx *out) {
    if (n_point == 1) {
        out[0].r = in[0].r;
        out[0].i = in[0].i;
        return;
    }
    kiss_fft_cpx *even_in =
        (kiss_fft_cpx *)KISS_FFT_MALLOC((n_point / 2) * sizeof(kiss_fft_cpx));
    kiss_fft_cpx *even_out =
        (kiss_fft_cpx *)KISS_FFT_MALLOC((n_point / 2) * sizeof(kiss_fft_cpx));
    kiss_fft_cpx *odd_in =
        (kiss_fft_cpx *)KISS_FFT_MALLOC((n_point / 2) * sizeof(kiss_fft_cpx));
    kiss_fft_cpx *odd_out =
        (kiss_fft_cpx *)KISS_FFT_MALLOC((n_point / 2) * sizeof(kiss_fft_cpx));

    for (int i = 0; i < n_point / 2; i++) {
        even_in[i] = in[i * 2];
        odd_in[i] = in[i * 2 + 1];
    }
    my_fft(n_point / 2, even_in, even_out);
    my_fft(n_point / 2, odd_in, odd_out);

    for (int i = 0; i < n_point / 2; i++) {
        float factor_r = cos(-2 * PI * i / n_point);
        float factor_i = sin(-2 * PI * i / n_point);

        out[i].r =
            even_out[i].r + odd_out[i].r * factor_r - odd_out[i].i * factor_i;
        out[i].i =
            even_out[i].i + odd_out[i].r * factor_i + odd_out[i].i * factor_r;
    }
    for (int i = n_point / 2; i < n_point; i++) {
        float factor_r = cos(-2 * PI * i / n_point);
        float factor_i = sin(-2 * PI * i / n_point);

        out[i].r = even_out[i - n_point / 2].r +
                   odd_out[i - n_point / 2].r * factor_r -
                   odd_out[i - n_point / 2].i * factor_i;
        out[i].i = even_out[i - n_point / 2].i +
                   odd_out[i - n_point / 2].r * factor_i +
                   odd_out[i - n_point / 2].i * factor_r;
    }
}

void reset(kiss_fft_cpx *data) {
    for (int i = 0; i < N; i++) {
        data[i].r = 0.0f;
        data[i].i = 0.0f;
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
    reset(out);
    TIMEIT(kiss_fft(cfg, in, out), 10);
    printf("1hz cos: %f\n", out[1].r / 64);
    printf("2hz sin: %f\n", -out[2].i / 64);
    printf("------\n");

    printf("dft manually:\n");
    reset(out);
    TIMEIT(my_dft(N, in, out), 10);
    printf("1hz cos: %f\n", out[1].r / 64);
    printf("2hz sin: %f\n", -out[2].i / 64);
    printf("------\n");

    printf("fft manually:\n");
    reset(out);
    TIMEIT(my_fft(N, in, out), 10);
    printf("1hz cos: %f\n", out[1].r / 64);
    printf("2hz sin: %f\n", -out[2].i / 64);
    printf("------\n");

    printf("ifft:\n");
    reset(in_restored);
    cfg = kiss_fft_alloc(N, 1, 0, 0);
    TIMEIT(kiss_fft(cfg, out, in_restored), 10);
    for (int i = 0; i < 10; i++) {
        printf("%f %f\n", in_restored[i].r / N, in[i].r);
    }
    printf("------\n");

    printf("dift manually:\n");
    reset(in_restored);
    TIMEIT(my_idft(N, out, in_restored), 10)
    for (int i = 0; i < 10; i++) {
        printf("%f %f\n", in_restored[i].r / N, in[i].r);
    }
}
