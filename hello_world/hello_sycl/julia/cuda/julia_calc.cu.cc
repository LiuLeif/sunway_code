#include "julia_calc.cu.h"

int HowManySteps(int x, int y, int width, int height, int zoom) {
  float cx = 0.285;
  float cy = 0.01;

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

extern void Julia(int height, int width, int zoom, void *data);

void JuliaCalculatorCu::Calc() {
  Julia(height_, width_, zoom_, data_);
}

JuliaCalculator* JuliaCalculator::get(size_t width, size_t height, void* data) {
  return new JuliaCalculatorCu(width, height, data);
}
