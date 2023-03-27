
#include <gcc-plugin.h>
#include <plugin-version.h>
#include <tree-core.h>
#include <tree-pretty-print.h>
#include <tree.h>

#include <iostream>

static void callback_finish_type(void *gcc_data, void *user_data) {
    std::cerr << "hello" << std::endl;
}

static void callback_finish_declaration(void *gcc_data, void *user_data) {
    std::cerr << "finish decl of: ";
    tree_node *tree = (tree_node *)gcc_data;
    debug_generic_expr(tree);
    std::cerr << "code: " << get_tree_code_name(TREE_CODE(tree)) << std::endl;
    debug_generic_expr(TREE_TYPE(tree));
}

static void callback_finish_unit(void *gcc_data, void *user_data) {}

static void callback_pre_genericize(void *gcc_data, void *user_data) {}

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
}
