#include <CL/sycl.hpp>
#include <iostream>
namespace sycl = cl::sycl;

class kernel_vector_add;
class kernel_vector_add_2;

void hello_usm_device() {
    // <<Setup host storage>>
    sycl::vec<float, 4> a = {1.0, 2.0, 3.0, 4.0};
    sycl::vec<float, 4> b = {1.0, 2.0, 3.0, 4.0};
    sycl::vec<float, 4> c = {0.0, 0.0, 0.0, 0.0};
    sycl::host_selector device;
    sycl::queue queue(device);
    std::cout << "Using "
              << queue.get_device().get_info<sycl::info::device::name>()
              << std::endl;

    float *device_a = sycl::malloc_device<float>(4, queue);
    float *device_b = sycl::malloc_device<float>(4, queue);
    float *device_c = sycl::malloc_device<float>(4, queue);

    queue.memcpy(device_a, &a, 16).wait();
    queue.memcpy(device_b, &b, 16).wait();

    queue.submit([&](sycl::handler &cgh) {
        cgh.single_task<class kernel_vector_add>([=]() {
            device_c[0] = device_a[0] + device_b[0];
            device_c[1] = device_a[1] + device_b[1];
            device_c[2] = device_a[2] + device_b[2];
            device_c[3] = device_a[3] + device_b[3];
        });
    });

    queue.wait();
    queue.memcpy(&c, device_c, 16).wait();
    // <<Print results>>
    std::cout << c.x() << "," << c.y() << "," << c.z() << "," << c.w()
              << std::endl;
}

void hello_usm_shared() {
    // <<Setup host storage>>
    sycl::vec<float, 4> a = {1.0, 2.0, 3.0, 4.0};
    sycl::vec<float, 4> b = {1.0, 2.0, 3.0, 4.0};
    sycl::host_selector device;
    sycl::queue queue(device);

    float *device_a = sycl::malloc_device<float>(4, queue);
    float *device_b = sycl::malloc_device<float>(4, queue);
    float *c = sycl::malloc_shared<float>(4, queue);

    queue.memcpy(device_a, &a, 16).wait();
    queue.memcpy(device_b, &b, 16).wait();

    queue.submit([&](sycl::handler &cgh) {
        cgh.single_task<class kernel_vector_add_2>([=]() {
            c[0] = device_a[0] + device_b[0];
            c[1] = device_a[1] + device_b[1];
            c[2] = device_a[2] + device_b[2];
            c[3] = device_a[3] + device_b[3];
        });
    });

    queue.wait();
    std::cout << c[0] << "," << c[1] << "," << c[2] << "," << c[3] << std::endl;
}

int main(int argc, char *argv[]) {
    hello_usm_device();
    hello_usm_shared();
    return 0;
}
