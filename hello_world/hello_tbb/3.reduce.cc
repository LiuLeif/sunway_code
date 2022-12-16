#include <tbb/blocked_range.h>
#include <tbb/parallel_for.h>
#include <tbb/tbb.h>

#include <iostream>

using namespace std;
using namespace tbb;

class Sum {
   public:
    int sum;
    void operator()(const blocked_range<int>& r) {
        for (int i = r.begin(); i != r.end(); ++i) {
            sum += i;
        }
    }
    void join(const Sum& y) { sum += y.sum; }
    Sum(Sum& x, split dummy) : sum(0) {}
    Sum() : sum(0) {}
};

#define N 1000
int main(int argc, char* argv[]) {
    cout << "-----" << endl;
    int sum = parallel_reduce(
        blocked_range<int>(0, N), 0,
        [=](const blocked_range<int>& r, int partial_sum) {
            for (int i = r.begin(); i != r.end(); i++) {
                partial_sum += i;
            }
            return partial_sum;
        },
        std::plus<int>());
    cout << sum << endl;

    cout << "-----" << endl;
    Sum x;
    parallel_reduce(blocked_range<int>(0, N), x);
    cout << x.sum << endl;
    return 0;
}
