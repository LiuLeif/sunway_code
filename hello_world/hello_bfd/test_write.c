// 2022-03-04 13:31
#include <assert.h>
#include <bfd.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int main(int argc, char* argv[]) {
    bfd* abfd = bfd_openw("/tmp/test.obj", NULL);
    assert(abfd != NULL);
    bfd_set_format(abfd, bfd_object);
    asymbol* symbol_table[2];
    memset(symbol_table, 0, sizeof(symbol_table));

    symbol_table[0] = bfd_make_empty_symbol(abfd);
    symbol_table[0]->name = "hello";
    symbol_table[0]->section = bfd_make_section_old_way(abfd, ".text");
    symbol_table[0]->flags = BSF_GLOBAL;
    symbol_table[0]->value = "0x1234";

    bfd_set_symtab(
        abfd, symbol_table, sizeof(symbol_table) / sizeof(symbol_table[0]) - 1);
    bfd_close(abfd);
}
