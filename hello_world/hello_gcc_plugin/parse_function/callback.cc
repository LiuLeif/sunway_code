// clang-format off
#include <gcc-plugin.h>
#include <plugin-version.h>
#include <print-tree.h>
#include <tree-core.h>
#include <tree.h>
#include <tree-iterator.h>
#include <tree-pretty-print.h>
// clang-format on

#include <iostream>

#include "plugin.h"

static void callback_finish_type(void *gcc_data, void *user_data) {
    std::cerr << "finish type of: ";
    tree_node *tree = (tree_node *)gcc_data;
    debug_generic_expr(tree);
}

static void callback_finish_declaration(void *gcc_data, void *user_data) {
    std::cerr << "======" << std::endl;
    std::cerr << "DECLARATION: " << std::endl;
    tree decl = (tree)gcc_data;
    // debug_tree(tree);
    std::cerr << "code: " << get_tree_code_name(TREE_CODE(decl)) << std::endl;
    std::cerr << "decl name: " << IDENTIFIER_POINTER(DECL_NAME(decl))
              << std::endl;

    tree tree_type = TREE_TYPE(decl);
    std::cerr << "tree_type code: " << get_tree_code_name(TREE_CODE(tree_type))
              << std::endl;
}

void traverse_function_body(tree t) {
    // t = build_int_cst(integer_type_node, 1);
#define CHECK_AND_CHANGE_CST(x)            \
    do {                                   \
        if (TREE_CODE(x) == INTEGER_CST) { \
            TREE_INT_CST_ELT(x, 0) = 1;    \
        }                                  \
    } while (0)

    if (t == NULL) {
        return;
    }
    if (TREE_CODE(t) == STATEMENT_LIST) {
        tree_stmt_iterator it = tsi_start(t);
        while (!tsi_end_p(it)) {
            traverse_function_body(tsi_stmt(it));
            tsi_next(&it);
        }
    } else if (TREE_CODE(t) == BIND_EXPR) {
        traverse_function_body(BIND_EXPR_BODY(t));
        traverse_function_body(BIND_EXPR_VARS(t));
    } else if (TREE_CODE(t) == VAR_DECL) {
        traverse_function_body(DECL_INITIAL(t));
        traverse_function_body(TREE_CHAIN(t));
        CHECK_AND_CHANGE_CST(DECL_INITIAL(t));
    } else if (EXPR_P(t)) {
        for (int i = 0; i < TREE_OPERAND_LENGTH(t); i++) {
            tree operand = TREE_OPERAND(t, i);
            if (EXPR_P(operand)) {
                traverse_function_body(operand);
            }
            CHECK_AND_CHANGE_CST(operand);
        }
    }
}

void callback_parse_function(void *event, void *__unused__) {
    std::cerr << "======" << std::endl;
    std::cerr << "FUNCTION: " << std::endl;
    tree t = (tree)event;
    printf(
        "%s %s\n", get_tree_code_name(TREE_CODE(t)),
        IDENTIFIER_POINTER(DECL_NAME(t)));
    debug_tree(DECL_SAVED_TREE(t));
    traverse_function_body(DECL_SAVED_TREE(t));
}

void register_callbacks(const char *base_name) {
    register_callback(
        base_name, PLUGIN_FINISH_TYPE, callback_finish_type,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_FINISH_DECL, callback_finish_declaration,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_FINISH_PARSE_FUNCTION, callback_parse_function,
        /* user_data */ NULL);
}
