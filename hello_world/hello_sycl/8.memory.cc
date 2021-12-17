#include <CL/sycl.hpp>
#include <iostream>
namespace sycl = cl::sycl;

class kernel_dummy;

int main(int argc, char* argv[]) {
    std::array<float, 10> a = {1.0, 2.0, 3.0, 4.0, 5.0,
                               6.0, 7.0, 8.0, 9.0, 10.0};
    sycl::gpu_selector device_selector;
    sycl::queue queue(device_selector, [](sycl::exception_list el) {
        for (auto ex : el) {
            try {
                std::rethrow_exception(ex);
            } catch (sycl::exception const& e) {
                std::cout << "Caught asynchronous SYCL exception:\n"
                          << e.what() << std::endl;
            }
        }
    });

    {
        sycl::buffer<float, 1> buff_a(
            a.data(), sycl::range<1>(10),
            {sycl::property::buffer::use_host_ptr()});

        queue.submit([&](sycl::handler& cgh) {
            // accessor 有 4 种 target:
            //
            // 1. global_buffer
            // 2. host_buffer
            // 3. consant_buffer
            // 4. local
            //
            // NOTE: 指定 cgh 参数时, target 默认是 global_buffer
            //
            // auto a_acc = buff_a.get_access<
            //     sycl::access::mode::read,
            //     sycl::access::target::global_buffer>( cgh);
            //
            // 除了 buffer.get_access, 还可以通过 accessor 构造函数:
            //
            // sycl::accessor<
            //     float, 1, sycl::access::mode::read,
            //     sycl::access::target::global_buffer>
            //     a_acc(buff_a, cgh);
            //
            // NOTE: 不指定 handle 时默认使用 host_buffer
            //
            // auto a_acc = buff_a.get_access<
            //     sycl::access::mode::read>();
            //
            // sycl::accessor<
            //     float, 1, sycl::access::mode::read,
            //     sycl::access::target::host_buffer>
            //     a_acc(buff_a);
            //
            // NOTE: constant_buffer 只支持 sycl::access::mode::read, 不支持
            // write: sycl::accessor<
            //     float, 1, sycl::access::mode::read,
            //     sycl::access::target::constant_buffer>
            //     const_acc(buff_a, cgh);
            //
            // NOTE: 如果使用 local_buffer, 因为 local_buffer 是 workgroup 内部
            // buffer, sycl 不支持通过初始数据来初始化, 只能指定一个大小
            //
            // sycl::accessor<
            //     float, 1, sycl::access::mode::read_write,
            //     sycl::access::target::local>
            //     const_acc(sycl::range<1>(10), cgh);
            //
            // NOTE: private memory 没有对应的 target, kernel
            // 中的局部变量会自动使用 private memory
            //
            // NOTE: 看起来 USM (malloc_{device,host,shared}) 没有 api 可以从
            // constant_buffer 或 local 分配内存

            sycl::accessor<
                float, 1, sycl::access::mode::read_write,
                sycl::access::target::global_buffer>
                global_acc(buff_a, cgh);

            sycl::accessor<
                float, 1, sycl::access::mode::read,
                sycl::access::target::constant_buffer>
                const_acc(buff_a, cgh);

            sycl::accessor<
                float, 1, sycl::access::mode::read_write,
                sycl::access::target::local>
                local_acc(sycl::range<1>(5), cgh);

            cgh.parallel_for<class kernel_dummy>(
                sycl::nd_range<1>(10, 5), [=](sycl::nd_item<1> item) {
                    // NOTE: malloc 或 malloc_device 均无法在 kernel 中使用
                    // float* x = (float*)malloc(4);
                    // sycl::malloc_device<float>(1, queue);
                    size_t local_id = item.get_local_linear_id();
                    size_t global_id = item.get_global_linear_id();
                    // 数组可以使用, 但依赖于动态内存分配的 stl 容器无法使用
                    std::array<float, 1> private_buf1 = {const_acc[global_id]};
                    float private_buf2[1] = {const_acc[global_id]};
                    global_acc[global_id] += private_buf1[0] + private_buf2[0];
                    // NOTE: accessor 还有一个 get_pointer 方法
                    auto ptr = global_acc.get_pointer();
                    auto const_ptr = const_acc.get_pointer();
                    auto local_ptr = local_acc.get_pointer();
                    *(ptr + global_id) += *(ptr + global_id);

                    printf(
                        "global ptr: %p, constant ptr: %p, local ptr: %p, "
                        "group: "
                        "%d\n",
                        ptr + global_id, const_ptr + global_id,
                        local_ptr + local_id, item.get_group_linear_id());
                });
        });

        sycl::accessor<
            float, 1, sycl::access::mode::read,
            sycl::access::target::host_buffer>
            host_acc(buff_a);

        auto host_ptr = host_acc.get_pointer();
        printf("host ptr: %p, host buffer: %p", host_ptr, a.data());

        printf("------\n");
        for (int i = 0; i < 10; i++) {
            printf("%f\n", *(host_ptr + i));
        }

        printf("------\n");
        for (int i = 0; i < 10; i++) {
            printf("%f\n", host_acc[i]);
        }

        printf("------\n");
        for (int i = 0; i < 10; i++) {
            printf("%f\n", a[i]);
        }
    }
    // NOTE: buff_a 析构, sycl 会把结果复制到 a
    printf("------\n");
    for (int i = 0; i < 10; i++) {
        printf("%f\n", a[i]);
    }

    return 0;
}
