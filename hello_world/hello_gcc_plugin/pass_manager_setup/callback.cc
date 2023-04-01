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
#include <context.h>

// clang-format on
const pass_data my_pass_data = {
    .type = GIMPLE_PASS,
    .name = "my_pass",
    .optinfo_flags = OPTGROUP_NONE,
    .tv_id = TV_NONE,
    .properties_required = PROP_gimple_any,
    .properties_provided = 0,
    .properties_destroyed = 0,
    .todo_flags_start = 0,
    .todo_flags_finish = 0};

struct my_pass : gimple_opt_pass {
   public:
    my_pass(gcc::context* ctxt) : gimple_opt_pass(my_pass_data, ctxt) {}

    virtual unsigned int execute(function* fun) override {
        printf("%s\n", function_name(fun));
        return 0;
    }

    virtual my_pass* clone() override { return this; }
};

struct register_pass_info my_passinfo {
    .pass = new my_pass(g), .reference_pass_name = "cfg",
    .ref_pass_instance_number = 1, .pos_op = PASS_POS_INSERT_AFTER
};

void register_callbacks(const char* base_name) {
    register_callback(base_name, PLUGIN_PASS_MANAGER_SETUP, NULL, &my_passinfo);
}
