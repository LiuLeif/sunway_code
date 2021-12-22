#ifndef JULIA_CALC_SYCL_H
#define JULIA_CALC_SYCL_H

#include <CL/sycl.hpp>
#include <iostream>

#include "../julia_calc.h"

namespace sycl = cl::sycl;

class JuliaCalculatorSycl : public JuliaCalculator {
  sycl::queue queue_;
  sycl::buffer<sycl::cl_uchar4, 2> img_;

 public:
  JuliaCalculatorSycl(size_t width, size_t height, void* data)
      : JuliaCalculator(width, height, data),
        queue_(sycl::host_selector{}),
        img_(sycl::range<2>(height, width)) {}

  void Calc();
};

#endif  // JULIA_CALC_SYCL_H
