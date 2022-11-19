// 2022-11-19 17:20
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int bisect(int* data, int N, int key) {
    int lo = 0;
    int hi = N;
    int mid = 0;
    while (lo < hi) {
        mid = (lo + hi) / 2;
        if (data[mid] == key) {
            return mid;
        }
        if (data[mid] > key) {
            hi = mid;
        } else {
            lo = mid + 1;
        }
    }
    return -1;
}

int main(int argc, char* argv[]) {
    int data[] = {1, 2, 3, 4, 5};
    printf("%d\n", bisect(data, sizeof(data) / sizeof(int), 1));
}
