#ifndef JULIA_CALC_H
#define JULIA_CALC_H

#include <iostream>

class JuliaCalculator {
 protected:
  size_t const width_;
  size_t const height_;
  int zoom_ = 1;
  void* data_;
  float cx_ = 0.285;
  float cy_ = 0.01;

 public:
  JuliaCalculator(size_t width, size_t height, void* data)
      : width_(width), height_(height), data_(data) {}

  virtual void Calc() = 0;

  void Zoom(float zoom) {
    zoom_ += int(zoom);
    if (zoom_ < 1) {
      zoom_ = 1;
    }
  }

  void SetC(float x, float y) {
    cx_ = x;
    cy_ = y;
  }

  static JuliaCalculator* get(size_t width, size_t height, void* data);
};

#endif  // JULIA_CALC_H
