// 2022-03-04 11:42
#include <assert.h>
#include <bfd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

void read_symtab(bfd *abfd) {
    int symtab_size = bfd_get_symtab_upper_bound(abfd);
    assert(symtab_size > 0);

    asymbol **symbol_table = (asymbol **)malloc(symtab_size);
    int num_symbols = bfd_canonicalize_symtab(abfd, symbol_table);
    printf("symtabl size: %d, num of symbols: %d\n", symtab_size, num_symbols);

    symbol_info info;
    for (int i = 0; i < num_symbols; i++) {
        if (symbol_table[i]->section == NULL) {
            continue;
        }
        bfd_symbol_info(symbol_table[i], &info);
        printf(
            "section: %s, symbol: %s -> 0x%lx, type: %x\n",
            symbol_table[i]->section->name, info.name, info.value, info.type);
    }
}

int main(int argc, char *argv[]) {
    bfd_init();
    /* bfd_openr 用来读 */
    /* bfd_openw 用来写 */
    bfd *abfd = bfd_openr("test.obj", NULL);
    assert(abfd != NULL);
    /* abfd 支持三种格式, 其中 elf 的 executable, o, so 都算 bfd_object */
    /*
     * if (bfd_check_format(abfd, bfd_archive)) {
     *     printf("found archive\n");
     * }
     * if (bfd_check_format(abfd, bfd_core)) {
     *     printf("found core\n");
     * }
     */
    assert(bfd_check_format(abfd, bfd_object) != 0);

    read_symtab(abfd);
    return 0;
}
