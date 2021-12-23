#include <stdint.h>
#include <stdio.h>

__device__ float2 complex_mul(float2 a, float2 b) {
    return {a.x * b.x - a.y * b.y, a.x * b.y + a.y * b.x};
}

__device__ float2 complex_add(float2 a, float2 b) {
    return {a.x + b.x, a.y + b.y};
}

__device__ float complex_norm(float2 a) { return a.x * a.x + a.y * a.y; }

__device__ int HowManySteps(float2 z, float2 c) {
    int MAX_ITERS = 255;
    float DIVERGENCE_LIMIT = 2.0;

    for (size_t i = MAX_ITERS; i > 0; i--) {
        z = complex_mul(z, z);
        z = complex_add(z, c);
        float norm = complex_norm(z);
        if (norm >= DIVERGENCE_LIMIT) {
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

    int count = HowManySteps(float2{zx, zy}, float2{cx, cy});
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
    // NOTE: 直接指定 kernel shape 为 (height, width) 不可行, 因为一个
    // block最多只 能有 1024 个 thread, 导致 width 不能超过 1024,
    // 这里是模拟了sycl 的 range 方 法
    JuliaKernel<<<ceil((height * width) / 32), 32>>>(
        height, width, zoom, dev_data, cx, cy, center_x, center_y);
    cudaDeviceSynchronize();
    cudaMemcpy(
        data, dev_data, sizeof(uchar4) * height * width,
        cudaMemcpyDeviceToHost);
}
