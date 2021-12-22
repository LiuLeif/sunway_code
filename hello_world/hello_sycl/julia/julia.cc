#include <cinder/CameraUi.h>
#include <cinder/TriMesh.h>
#include <cinder/app/App.h>
#include <cinder/app/RendererGl.h>
#include <cinder/gl/gl.h>

#include <iostream>

#include "julia_calc.h"

constexpr size_t WIDTH = 1920;
constexpr size_t HEIGHT = 1080;

class JuliaApp : public ci::app::App {
  // Texture for displaying the set
  ci::gl::Texture2dRef tex_;
  void* data_;
  JuliaCalculator* calc_;

 public:
  JuliaApp()
      : data_(malloc(WIDTH * HEIGHT * 4)),
        calc_(JuliaCalculator::get(WIDTH, HEIGHT, data_)) {}

  void setup() override {
    this->tex_ = ci::gl::Texture2d::create(
        nullptr, GL_RGBA, WIDTH, HEIGHT,
        ci::gl::Texture2d::Format()
            .dataType(GL_UNSIGNED_BYTE)
            .internalFormat(GL_RGBA));
  }

  void update() override {
    printf("update\n");
    calc_->Calc();
  }

  void draw() override {
    ci::gl::clear();
    tex_->update(data_, GL_RGBA, GL_UNSIGNED_BYTE, 0, WIDTH, HEIGHT);
    ci::Rectf screen(0, 0, getWindow()->getWidth(), getWindow()->getHeight());
    ci::gl::draw(tex_, screen);
  }

  void mouseWheel(ci::app::MouseEvent event) override {
    auto inc = event.getWheelIncrement();
    printf("mouse whell %f\n", inc);
    calc_->Zoom(inc);
  }

  void mouseDrag(ci::app::MouseEvent event) override {
    auto x = event.getX() / double(WIDTH);
    auto y = event.getY() / double(HEIGHT);
  }
};

CINDER_APP(JuliaApp, ci::app::RendererGl(ci::app::RendererGl::Options{}))
