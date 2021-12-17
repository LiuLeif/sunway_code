#include <CL/sycl.hpp>
#include <iostream>
namespace sycl = cl::sycl;

class kernel_dummy_1;
class kernel_dummy_2;
class kernel_dummy_3;

int main(int argc, char* argv[]) {
    // A=a*2
    // B=b*3
    // C=A+B
    int32_t a = 1;
    int32_t b = 2;
    int32_t A = 0;
    int32_t B = 0;
    int32_t C = 0;

    sycl::queue queue(sycl::gpu_selector{});

    sycl::buffer<int32_t, 1> buff_a(&a, sycl::range<1>(1));
    sycl::buffer<int32_t, 1> buff_b(&b, sycl::range<1>(1));
    sycl::buffer<int32_t, 1> buff_A(&A, sycl::range<1>(1));
    sycl::buffer<int32_t, 1> buff_B(&B, sycl::range<1>(1));
    sycl::buffer<int32_t, 1> buff_C(&C, sycl::range<1>(1));
    // NOTE: 在 cuda 中, 同一个 stream 里的 kernel 都是串行执行的, 但按照
    // sycl的文 档, kernel 被 submit 的顺序是用户期望的 kernel 执行的顺序 (与
    // cuda 相同), 但 是 accessor 决定的数据依赖关系可以使没有数据依赖的 kernel
    // 可以并行执行或改变执行顺序.
    //
    // 需要注意的是 sycl runtime 并不能完全自己推导出来 DAG, 例如, 若 submit 的
    // kernel 顺序变成 (C=A+B, A=2a, B=2B), 则 C 的结果会是错误的 0 (而不是 8)
    //
    // 原因在于, 即使多个 kernel 中有些是 `read A`, 有些是 `write A`, sycl
    // runtime 也并不能推导出 `read A` 依赖 `write A`, 需要用户通过 submit
    // 来告诉 runtime 针对 `A` 的 data race 是 read-write 还是 write-read.
    //
    // 这一点与普通 cpu 的指令重排非常类似:
    //
    // 1. x=a+1
    // 2. y=b+1
    // 3. z=x+y
    //
    // 编译器通过分析代码能知道 (1,3),(2,3) 有 data race (都是
    // read-after-write), 但 (1,2) 并没有, 所以 1,2可以自由的乱序执行
    queue.submit([&](sycl::handler& cgh) {
        auto b_acc = buff_b.get_access<sycl::access::mode::read>(cgh);
        auto B_acc = buff_B.get_access<sycl::access::mode::discard_write>(cgh);

        cgh.single_task<class kernel_dummy_1>([=]() {
            printf("B=b*3\n");
            B_acc[0] = b_acc[0] * 3;
        });
    });

    queue.submit([&](sycl::handler& cgh) {
        auto a_acc = buff_a.get_access<sycl::access::mode::read>(cgh);
        auto A_acc = buff_A.get_access<sycl::access::mode::discard_write>(cgh);

        cgh.single_task<class kernel_dummy_3>([=]() {
            printf("A=a*2\n");
            A_acc[0] = a_acc[0] * 2;
        });
    });

    queue.submit([&](sycl::handler& cgh) {
        auto A_acc = buff_A.get_access<sycl::access::mode::read>(cgh);
        auto B_acc = buff_B.get_access<sycl::access::mode::read>(cgh);
        auto C_acc = buff_C.get_access<sycl::access::mode::discard_write>(cgh);

        cgh.single_task<class kernel_dummy_2>([=]() {
            printf("C=A+B\n");
            C_acc[0] = A_acc[0] + B_acc[0];
        });
    });

    auto host_acc = buff_C.get_access<sycl::access::mode::read>();
    printf("%d\n", host_acc[0]);
    return 0;
}
