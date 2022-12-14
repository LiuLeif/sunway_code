#include <tbb/blocked_range.h>
#include <tbb/parallel_for.h>
#include <tbb/tbb.h>

#include <iostream>

using namespace std;
int main(int argc, char *argv[]) {
    tbb::parallel_for(
        tbb::blocked_range<int>(0, 10), [=](const tbb::blocked_range<int> &r) {
            for (int i = r.begin(); i != r.end(); i++) {
                cout << i << endl;
            }
        });
    return 0;
}
