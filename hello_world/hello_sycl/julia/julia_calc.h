#ifndef JULIA_CALC_H
#define JULIA_CALC_H

#include <CL/sycl.hpp>
#include <iostream>
namespace sycl = cl::sycl;

class JuliaCalculator {
    // Dimensions of the image to be calculated
    size_t const width_;
    size_t const height_;
    int zoom_ = 1;

    sycl::queue queue_;
    sycl::buffer<sycl::cl_uchar4, 2> img_;
    void* data_;

   public:
    JuliaCalculator(size_t width, size_t height, void* data)
        : width_(width),
          height_(height),
          data_(data),
          queue_(sycl::host_selector{}),
          img_(sycl::range<2>(height, width)) {}

    void Calc();
};

#endif  // JULIA_CALC_H
