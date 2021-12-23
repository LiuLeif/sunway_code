#include <stdint.h>
#include <stdio.h>

// NOTE: 这里误指定 zx, zy 为 int 竟然能编译通过...导致结果错误
__device__ int HowManySteps(float zx, float zy, float cx, float cy) {
  float zx2 = 0.0;
  float zy2 = 0.0;
  float abs_sq = 0.0;

  int MAX_ITERS = 255;
  float DIVERGENCE_LIMIT = 2.0;

  for (size_t i = MAX_ITERS; i > 0; i--) {
    zx2 = zx * zx - zy * zy + cx;
    zy2 = 2.0 * zx * zy + cy;

    zx = zx2;
    zy = zy2;

    abs_sq = zx * zx + zy * zy;

    if (abs_sq >= DIVERGENCE_LIMIT) {
      return i;
    }
  }

  return 0;
}

__global__ void JuliaKernel(
    int height, int width, int zoom, uchar4 *dev_data, float cx, float cy) {
  int x = threadIdx.x;
  int y = blockIdx.x;

  float zx = (x - 0.5 * width) / (0.5 * width * zoom);
  float zy = (y - 0.5 * height) / (0.5 * height * zoom);

  int count = HowManySteps(zx, zy, cx, cy);
  int color = (count << 21) + (count << 10) + (count << 3);
  dev_data[x * height + y] = {
      (uint8_t)(color >> 16), (uint8_t)(color >> 8), (uint8_t)color,
      (uint8_t)255};
}

void Julia(int height, int width, int zoom, void *data, float cx, float cy) {
  static uchar4 *dev_data = 0;
  if (dev_data == 0) {
    cudaMalloc(&dev_data, sizeof(uchar4) * height * width);
  }
  JuliaKernel<<<height, width>>>(height, width, zoom, dev_data, cx, cy);
  // cudaError_t err = cudaGetLastError();
  // if (err != cudaSuccess) {
  //   printf("CUDA Error: %s\n", cudaGetErrorString(err));
  // }
  cudaDeviceSynchronize();
  cudaMemcpy(
      data, dev_data, sizeof(uchar4) * height * width, cudaMemcpyDeviceToHost);
}
