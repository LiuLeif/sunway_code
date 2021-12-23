#include "julia_calc.h"

extern void Julia(
    int height, int width, float zoom, void* data, float cx, float cy,
    float center_x, float center_y);

void JuliaCalculatorCu::Calc() {
    Julia(height_, width_, zoom_, data_, cx_, cy_, center_x_, center_y_);
}

JuliaCalculator* JuliaCalculator::get(size_t width, size_t height, void* data) {
    return new JuliaCalculatorCu(width, height, data);
}
