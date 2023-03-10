// 2023-03-10 14:42
// gcc test.c printf.c -nostdlib
extern int printf(const char *, ...);
int main(int argc, char *argv[]) {
    /* 使用 hello\n 时链接时找不到 puts */
    printf("hello\n");
    /* 使用 hello 时链接时不会去找 puts */
    printf("hello");
}

/* 这个问题是因为 gcc 的优化导致的: gcc 认为 printf 有众所周知的意义, 当要打印的
 * string 不包含 %, 或者包含 \n (即不需要缓冲), 则 gcc 认为可以用 puts 来代替 printf
 * 通过 -fno-builtin 可以禁止这个优化
 *  */
