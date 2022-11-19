// 2022-11-19 15:36
#include <stdio.h>

static inline void swap(int* data, int a, int b) {
    int tmp = data[b];
    data[b] = data[a];
    data[a] = tmp;
}

void quick_sort(int* data, int from, int to) {
    if (from >= to) {
        return;
    }
    int pivot = from;
    int curr = from + 1;
    while (curr < to) {
        if (data[curr] >= data[pivot]) {
            curr += 1;
        } else {
            swap(data, curr, pivot + 1);
            swap(data, pivot, pivot + 1);
            pivot += 1;
            curr += 1;
        }
    }
    quick_sort(data, from, pivot);
    quick_sort(data, pivot + 1, to);
}

int main(int argc, char* argv[]) {
    int data[] = {0, 1, 2, 1, 2};
    int N = sizeof(data) / sizeof(int);
    quick_sort(data, 0, N);
    for (int i = 0; i < N; i++) {
        printf("%d \n", data[i]);
    }
}
