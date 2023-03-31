// clang-format off
#include <gcc-plugin.h>
#include <plugin-version.h>
#include <print-tree.h>
#include <tree-core.h>
#include <tree.h>
#include <tree-iterator.h>
#include <tree-pretty-print.h>
#include <stringpool.h>
#include <cgraph.h>
#include <gimplify.h>
#include <tree-pass.h>
#include <iostream>
// clang-format on

extern tree pushdecl(tree x);

// NOTE: 这个 plugin 给生成全局变量生成 getter 函数, 例如 __get_a

static void callback_finish_decl(void *gcc_data, void *user_data) {
    tree t = (tree)gcc_data;

    if (TREE_CODE(t) != VAR_DECL) {
        return;
    }
    if (DECL_CONTEXT(t) != NULL) {
        return;
    }
    if (! SCALAR_FLOAT_TYPE_P(TREE_TYPE(t)) && ! INTEGRAL_TYPE_P(TREE_TYPE(t))) {
        return;
    }
    // debug_tree(t);
    printf("found global var: %s\n", IDENTIFIER_POINTER(DECL_NAME(t)));

    tree decl_type = TREE_TYPE(t);
    tree ftype = build_function_type_list(decl_type, NULL_TREE);
    const char *var_name = IDENTIFIER_POINTER(DECL_NAME(t));

    tree decl = build_fn_decl(("__get_" + std::string(var_name)).c_str(), ftype);
    tree result =
        build_decl(DECL_SOURCE_LOCATION(t), RESULT_DECL, 0, decl_type);

    DECL_EXTERNAL(decl) = 0;
    DECL_INITIAL(decl) = make_node(BLOCK);
    DECL_RESULT(decl) = result;

    pushdecl(decl);

    tree modify = build2(MODIFY_EXPR, TREE_TYPE(result), result, t);
    modify = build1(RETURN_EXPR, void_type_node, modify);
    tree bind =
        build3(BIND_EXPR, void_type_node, NULL_TREE, modify, make_node(BLOCK));
    DECL_SAVED_TREE(decl) = bind;

    gimplify_function_tree(decl);
    cgraph_node::add_new_function(decl, false);
}

void register_callbacks(const char *base_name) {
    register_callback(
        base_name, PLUGIN_FINISH_DECL, callback_finish_decl, NULL);
}
