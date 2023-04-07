#include <stdio.h>

void trace() { printf("trace\n"); }

void foo(char* x) { int a = 0; }

char a[] = "hello";
int main(int argc, char* argv[]) {
    foo("a");
    return 0;
}
