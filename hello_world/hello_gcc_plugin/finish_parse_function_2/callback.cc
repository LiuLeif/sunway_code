// clang-format off
#include <gcc-plugin.h>
#include <plugin-version.h>
#include <print-tree.h>
#include <tree-core.h>
#include <tree.h>
#include <tree-iterator.h>
#include <tree-pretty-print.h>
#include <iostream>
// clang-format on

// NOTE: 这个 plugin 会把所有的 int 常量变成 1

void callback_parse_function(void *data, void *__unused__) {
    std::cerr << "======" << std::endl;
    std::cerr << "FUNCTION: " << std::endl;
    tree t = (tree)data;
    // debug_tree(t);
    const char *fname = IDENTIFIER_POINTER(DECL_NAME(t));
    if (strcmp(fname, "trace") == 0) {
        return;
    }
    printf("%s %s\n", get_tree_code_name(TREE_CODE(t)), fname);

    t = DECL_SAVED_TREE(t);
    tree ftype = build_function_type_list(
        void_type_node, build_pointer_type(char_type_node), NULL_TREE);
    tree fndecl = build_fn_decl("trace", ftype);

    if (TREE_CODE(t) == STATEMENT_LIST) {
        // NOTE: it seems `main` starts wil `statement_list`, while other
        // function starts with `bind_expr`
        tree_stmt_iterator iter = tsi_start(t);
        t = tsi_stmt(iter);
    }

    tree body = BIND_EXPR_BODY(t);
    tree_stmt_iterator iter = tsi_start(body);
    tree call_stmt = build_call_expr(
        fndecl, 1, build_string_literal(strlen(fname) + 1, fname));
    tsi_link_before(&iter, call_stmt, TSI_SAME_STMT);
}

void register_callbacks(const char *base_name) {
    register_callback(
        base_name, PLUGIN_FINISH_PARSE_FUNCTION, callback_parse_function, NULL);
}
