#include <stdint.h>

__device__ int HowManySteps(int zx, int zy) {
  float cx = 0.285;
  float cy = 0.01;

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

__global__ void JuliaKernel(int height, int width, int zoom, uchar4 *dev_data) {
  int x = threadIdx.x;
  int y = blockIdx.x;

  float zx = (x - 0.5 * width) / (0.5 * width * zoom);
  float zy = (y - 0.5 * height) / (0.5 * height * zoom);

  int count = HowManySteps(zx, zy);
  int color = (count << 21) + (count << 10) + (count << 3);
  dev_data[x * height + y] = {
      (uint8_t)(color >> 16), (uint8_t)(color >> 8), (uint8_t)color,
      (uint8_t)255};
}

void Julia(int height, int width, int zoom, void *data) {
  uchar4 *dev_data;
  cudaMalloc(&dev_data, sizeof(uchar4) * height * width);
  JuliaKernel<<<height, width>>>(height, width, zoom, dev_data);
  cudaMemcpy(
      data, dev_data, sizeof(uchar4) * height * width, cudaMemcpyDeviceToHost);
}
