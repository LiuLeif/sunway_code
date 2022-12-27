// 2022-12-27 14:48
#include <mpi.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    int pid, np;
    MPI_Init(&argc, &argv);
    MPI_Comm_rank(MPI_COMM_WORLD, &pid);
    MPI_Comm_size(MPI_COMM_WORLD, &np);
    int number[12] = {0};
    if (pid == 0) {
        for (int i = 0; i < 12; i++) {
            number[i] = i;
        }
    }

    int sub_number[3] = {0};
    MPI_Scatter(number, 3, MPI_INT, sub_number, 3, MPI_INT, 0, MPI_COMM_WORLD);

    int local_sum = 0;
    for (int i = 0; i < 3; i++) {
        local_sum += sub_number[i];
    }
    int sum = 0;
    MPI_Reduce(&local_sum, &sum, 1, MPI_INT, MPI_SUM, 0, MPI_COMM_WORLD);
    if (pid == 0) {
        printf("%d\n", sum);
    }

    MPI_Allreduce(&local_sum, &sum, 1, MPI_INT, MPI_SUM, MPI_COMM_WORLD);
    MPI_Barrier(MPI_COMM_WORLD);
    printf("pid: %d, %d\n", pid, sum);
    MPI_Finalize();
    return 0;
}
