#include <stdint.h>
#include <stdio.h>

// NOTE: 这里误指定 zx, zy 为 int 也能编译通过...导致结果错误
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
    int height, int width, float zoom, uchar4 *dev_data, float cx, float cy,
    float center_x, float center_y) {
  int global_id = blockIdx.x * blockDim.x + threadIdx.x;
  int x = (int)(global_id / height);
  int y = global_id - x * height;

  float zx = (x - 0.5 * width) / (0.5 * width * zoom) + center_x;
  float zy = (y - 0.5 * height) / (0.5 * height * zoom) + center_y;

  int count = HowManySteps(zx, zy, cx, cy);
  int color = (count << 21) + (count << 10) + (count << 3);
  dev_data[x * height + y] = {
      (uint8_t)(color >> 16), (uint8_t)(color >> 8), (uint8_t)color,
      (uint8_t)255};
}

void Julia(
    int height, int width, float zoom, void *data, float cx, float cy,
    float center_x, float center_y) {
  static uchar4 *dev_data = 0;
  if (dev_data == 0) {
    cudaMalloc(&dev_data, sizeof(uchar4) * height * width);
  }
  // NOTE: 直接指定 kernel shape 为 (height, width) 不可行, 因为一个 block
  // 最多只能有 1024 个 thread, 导致 width 不能超过 1024, 这里实现上是模拟了
  // sycl 的 range 方法
  JuliaKernel<<<ceil((height * width) / 32), 32>>>(
      height, width, zoom, dev_data, cx, cy, center_x, center_y);
  // cudaError_t err = cudaGetLastError();
  // if (err != cudaSuccess) {
  //   printf("CUDA Error: %s\n", cudaGetErrorString(err));
  // }
  cudaDeviceSynchronize();
  cudaMemcpy(
      data, dev_data, sizeof(uchar4) * height * width, cudaMemcpyDeviceToHost);
}
