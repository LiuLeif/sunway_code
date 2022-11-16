#include <Eigen/Dense>
#include <iostream>

using Eigen::Array22d;
using Eigen::MatrixXd;
using Eigen::Vector4f;

// void test_matrix() {
//     Array22d m1;
//     m1 << 1, 2, 3, 4;

//     Array22d m2;
//     m2 << 1, 2, 3, 4;

//     Array22d m3 = m1 + m2;
//     std::cout << m3 << std::endl;
// }
int main() {
    Vector4f a, b;
    a << 1, 2, 3, 4;
    b << 1, 2, 3, 4;
    a += b;
    // test_matrix();
    std::cout << a << std::endl;
}
