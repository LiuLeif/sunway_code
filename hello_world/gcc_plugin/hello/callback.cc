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
    tree_node *tree = (tree_node *)gcc_data;

    std::cerr << "code: " << get_tree_code_name(TREE_CODE(tree)) << std::endl;
    std::cerr << "decl name: " << IDENTIFIER_POINTER(DECL_NAME(tree))
              << std::endl;

    tree_node *tree_type = TREE_TYPE(tree);
    std::cerr << "tree_type code: " << get_tree_code_name(TREE_CODE(tree_type))
              << std::endl;
}

static void callback_finish_unit(void *gcc_data, void *user_data) {}

static void callback_pre_genericize(void *gcc_data, void *user_data) {}

void check_args(tree t) {
    for (int i = 0; i < TREE_OPERAND_LENGTH(t); i++) {
        tree operand = TREE_OPERAND(t, i);
        if (TREE_CODE(operand) == INTEGER_CST) {
            TREE_OPERAND(t, i) = build_int_cst(integer_type_node, 10);
        }
        if (EXPR_P(operand)) {
            check_args(operand);
        }
    }
}
void traverse_function_body(tree t, int level) {
    // printf("%s\n", get_tree_code_name(TREE_CODE(t)));
    switch (TREE_CODE(t)) {
        case STATEMENT_LIST: {
            tree_stmt_iterator it = tsi_start(t);
            while (!tsi_end_p(it)) {
                traverse_function_body(tsi_stmt(it), level + 4);
                tsi_next(&it);
            }
            break;
        }
        case BIND_EXPR: {
            tree body = BIND_EXPR_BODY(t);
            traverse_function_body(body, level + 4);
            check_args(body);
            // tree modify = TREE_OPERAND(body, 0);
            // tree plus = TREE_OPERAND(modify, 1);
            // debug_tree(body);
            // TREE_OPERAND(plus, 1) = build_int_cst(integer_type_node, 11);
            break;
        }
        default:
            break;
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
    traverse_function_body(DECL_SAVED_TREE(t), 0);
}

static void callback_finish(void *gcc_data, void *user_data) {}

static void callback_register_attribute(void *gcc_data, void *user_data) {}

static void callback_start_unit(void *gcc_data, void *user_data) {}

static void callback_registering_pragmas(void *gcc_data, void *user_data) {}

static void callback_all_passes_start(void *gcc_data, void *user_data) {}

static void callback_all_passes_end(void *gcc_data, void *user_data) {}

static void callback_pass_execution(void *gcc_data, void *user_data) {}

void register_callbacks(const char *base_name) {
    register_callback(
        base_name, PLUGIN_FINISH_TYPE, callback_finish_type,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_FINISH_DECL, callback_finish_declaration,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_FINISH_UNIT, callback_finish_unit,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_PRE_GENERICIZE, callback_pre_genericize,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_FINISH, callback_finish,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_ATTRIBUTES, callback_register_attribute,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_START_UNIT, callback_start_unit,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_PRAGMAS, callback_registering_pragmas,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_ALL_PASSES_START, callback_all_passes_start,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_ALL_PASSES_END, callback_all_passes_end,
        /* user_data */ NULL);

    register_callback(
        base_name, PLUGIN_PASS_EXECUTION, callback_pass_execution,
        /* user_data */ NULL);
    register_callback(
        base_name, PLUGIN_FINISH_PARSE_FUNCTION, callback_parse_function,
        /* user_data */ NULL);
}
