#include "julia_calc.h"

#include <CL/sycl.hpp>
#include <iostream>
namespace sycl = cl::sycl;

int HowManySteps(
    int x, int y, int width, int height, float zoom, float cx, float cy) {
    float zx = (x - 0.5 * width) / (0.5 * width * zoom);
    float zy = (y - 0.5 * height) / (0.5 * height * zoom);

    float zx2 = 0.0;
    float zy2 = 0.0;
    float abs_sq = 0.0;

    static constexpr size_t MAX_ITERS = 255;
    static constexpr float DIVERGENCE_LIMIT = 2.0;

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

void JuliaCalculatorSycl::Calc() {
    queue_.submit([&](sycl::handler& cgh) {
        auto img_acc = img_.get_access<sycl::access::mode::read_write>(cgh);

        // NOTE: 这样写是因为 kernel 无法 capture this...
        int width = width_;
        int height = height_;
        float zoom = zoom_;
        float cx = cx_;
        float cy = cy_;
        cgh.parallel_for<class JuliaCalculator>(
            sycl::range<2>(height, width), [=](sycl::item<2> item) {
                int count = HowManySteps(
                    item.get_id(1), item.get_id(0), width, height, zoom, cx,
                    cy);
                int color = (count << 21) + (count << 10) + (count << 3);
                img_acc[item] = {
                    (uint8_t)(color >> 16), (uint8_t)(color >> 8),
                    (uint8_t)color, (uint8_t)255};
            });
        cgh.copy(img_acc, data_);
    });
    // auto host_acc = img_.get_access<sycl::access::mode::read>();
    // memcpy(data_, host_acc.get_pointer(), width_ * height_ * 4);
    queue_.wait();
}

JuliaCalculator* JuliaCalculator::get(size_t width, size_t height, void* data) {
    return new JuliaCalculatorSycl(width, height, data);
}
