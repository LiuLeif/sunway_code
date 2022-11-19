// 2022-11-19 22:47
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void merge_sort(int *data, int N) {
    if (N <= 1) {
        return;
    }
    int mid = N / 2;
    int *left = data;
    int *right = data + mid;
    merge_sort(left, mid);
    merge_sort(right, N - mid);

    int *tmp = (int *)malloc(N * sizeof(int));
    int li = 0;
    int ri = 0;
    for (int i = 0; i < N; i++) {
        int min = 0;
        if (li >= mid) {
            min = right[ri++];
        } else if (ri >= N - mid) {
            min = left[li++];
        } else {
            int a = right[ri];
            int b = left[li];
            if (a < b) {
                min = a;
                ri++;
            } else {
                min = b;
                li++;
            }
        }
        tmp[i] = min;
    }
    memcpy(data, tmp, N * sizeof(int));
    free(tmp);
}
int main(int argc, char *argv[]) {
    int data[] = {2, 1, 2, 3, 4, -1, 0, 0};
    int N = sizeof(data) / sizeof(int);
    merge_sort(data, N);
    for (int i = 0; i < N; i++) {
        printf("%d ", data[i]);
    }
    printf("\n");
}
