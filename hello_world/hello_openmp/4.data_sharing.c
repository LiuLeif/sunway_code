// 2022-12-05 19:08
#include <omp.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void test_private() {
    printf("---%s--- \n", __FUNCTION__);
    int sum = 0;
#pragma omp parallel for private(sum) num_threads(2)
    for (int i = 0; i < 10; i++) {
        sum += 1;
        printf("%d %d\n", sum, omp_get_thread_num());
    }
    printf("------\n");
    sum = 0;
#pragma omp parallel for num_threads(2)
    for (int i = 0; i < 10; i++) {
        /* NOTE: 这里缺少同步, sum += 1 会出现 race */
        sum += 1;
        printf("%d %d\n", sum, omp_get_thread_num());
    }
    printf("------\n");
    sum = 0;
#pragma omp parallel num_threads(2)
    {
        int private_sum = sum;
#pragma omp for
        for (int i = 0; i < 10; i++) {
            private_sum += 1;
            printf("%d %d\n", private_sum, omp_get_thread_num());
        }
    }
}

int main(int argc, char *argv[]) { test_private(); }
