#include "julia_calc.h"

extern void Julia(
    int height, int width, int zoom, void* data, float cx, float cy);

void JuliaCalculatorCu::Calc() {
  Julia(height_, width_, zoom_, data_, cx_, cy_);
}

JuliaCalculator* JuliaCalculator::get(size_t width, size_t height, void* data) {
  return new JuliaCalculatorCu(width, height, data);
}
