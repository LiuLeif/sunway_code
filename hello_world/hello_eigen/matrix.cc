#include <Eigen/Dense>
#include <iostream>

using Eigen::Vector4f;

int main() {
    Eigen::Array4f a;
    a << 1, 2, 3, 4.1;

    std::cout << a.round() << std::endl;
}
