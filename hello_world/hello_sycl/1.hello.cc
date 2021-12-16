#include <CL/sycl.hpp>
#include <iostream>
namespace sycl = cl::sycl;

class kernel_vector_add;
class kernel_vector_add_2;
int main(int argc, char* argv[]) {
    // <<Setup host storage>>
    // float4 是 sycl::vec<float,4> 的别名
    sycl::float4 a = {1.0, 2.0, 3.0, 4.0};
    sycl::float4 b = {1.0, 2.0, 3.0, 4.0};
    sycl::float4 c = {0.0, 0.0, 0.0, 0.0};
    // <<Initialize device selector>>
    sycl::host_selector device_selector;
    // <<Initialize queue>>
    sycl::queue queue(device_selector);
    {
        //   <<Setup device storage>>
        //   sycl::range<dims>(dim1,dim2,...)
        sycl::buffer<sycl::float4, 1> buff_a(&a, sycl::range<1>(1));
        sycl::buffer<sycl::float4, 1> buff_b(&b, sycl::range<1>(1));
        sycl::buffer<sycl::float4, 1> buff_c(&c, sycl::range<1>(1));
        //   <<Execute kernel>>
        queue.submit([&](sycl::handler& cgh) {
            auto a_acc = buff_a.get_access<sycl::access::mode::read>(cgh);
            auto b_acc = buff_b.get_access<sycl::access::mode::read>(cgh);
            auto c_acc =
                buff_c.get_access<sycl::access::mode::discard_write>(cgh);

            // single_task 表示只启动一个线程
            // 更常用的是 parallel_for 以启动多个线程
            // [=] 表示 lambda 使用 copy 来 capture 自由变量
            // [&] 表示 lambda 使用 refernece 来 capture 自由变量
            // sycl 的 kernel 只支持 [=] 而不支持 [&]
            cgh.single_task<class kernel_vector_add>(
                [=]() { c_acc[0] = a_acc[0] + b_acc[0]; });
        });

        // NOTE: queue.submit 后是异步的, 这里直接访问还没有结果.
        // std::cout << c.x() << "," << c.y() << "," << c.z() << "," << c.w()
        //           << std::endl;
        //
        // 原始的 buffer `c` 需要等待异步执行完毕并且 device copy
        // 回来才能拿到结果
        //
        // 有两种方式可以用来等待异常执行的结果:
        //
        // 1. 通过 c++ 作用域, 当 buffer_c 离开作用域是, sycl 会强制一个
        // `等待结果` 和 `device copy` 的动作
        // 2. 通过 host accessor 表示对数据的依赖
    }

    // <<Print results>>
    std::cout << c.x() << "," << c.y() << "," << c.z() << "," << c.w()
              << std::endl;

    // NOTE: 通过 host accessor 表示数据的依赖及读取数据
    c = {0.0, 0.0, 0.0, 0.0};
    sycl::buffer<sycl::float4, 1> buff_a(&a, sycl::range<1>(1));
    sycl::buffer<sycl::float4, 1> buff_b(&b, sycl::range<1>(1));
    sycl::buffer<sycl::float4, 1> buff_c(&c, sycl::range<1>(1));
    queue.submit([&](sycl::handler& cgh) {
        auto a_acc = buff_a.get_access<sycl::access::mode::read>(cgh);
        auto b_acc = buff_b.get_access<sycl::access::mode::read>(cgh);
        auto c_acc = buff_c.get_access<sycl::access::mode::discard_write>(cgh);

        cgh.single_task<class kernel_vector_add_2>(
            // c_acc[0] 是因为 float4 类型只有一个元素 (虽然它包含 4 个数)
            [=]() { c_acc[0] = a_acc[0] + b_acc[0]; });
    });
    // NOTE: c_acc 表示的对 buff_c 的依赖会要求 sycl 在这里有一个 wait,
    // 等等异步操作的结束
    auto c_acc = buff_c.get_access<sycl::access::mode::read>();
    std::cout << c_acc[0].x() << "," << c_acc[0].y() << "," << c_acc[0].z()
              << "," << c_acc[0].w() << std::endl;
    // NOTE: 由于 buff_c 还有效, 所以 `c` 会由 buffer_c 接管, 对 `c`
    // 的直接访问无效
    // std::cout << c.x() << "," << c.y() << "," << c.z() << ","
    // << c.w() << std::endl;
    return 0;
}
