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
#include "basic-block.h"
#include <gimple.h>
#include <gimple-iterator.h>
#include <gimple-pretty-print.h>
#include <context.h>
#include <ssa.h>
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
    .todo_flags_finish = 0,
};

struct my_pass : gimple_opt_pass {
   public:
    my_pass(gcc::context* ctxt) : gimple_opt_pass(my_pass_data, ctxt) {}

    virtual void mark_stmt_visited(gimple* stmt) const {
        if (gimple_visited_p(stmt)) {
            return;
        }
        gimple_set_visited(stmt, true);
        use_operand_p use_p;
        ssa_op_iter i;
        FOR_EACH_SSA_USE_OPERAND(use_p, stmt, i, SSA_OP_USE) {
            tree use = USE_FROM_PTR(use_p);
            gimple* def = SSA_NAME_DEF_STMT(use);
            mark_stmt_visited(def);
        }
    }
    virtual unsigned int execute(function* fun) override {
        printf("===FNDECL===\n");
        printf("%s\n", function_name(fun));

        tree var;
        int i;
        FOR_EACH_VEC_SAFE_ELT(cfun->local_decls, i, var) {}

        basic_block bb;
        gimple_stmt_iterator gsi;

        printf("===BODY===\n");
        FOR_EACH_BB_FN(bb, fun)
        for (gsi = gsi_start_bb(bb); !gsi_end_p(gsi); gsi_next(&gsi)) {
            gimple* stmt = gsi_stmt(gsi);
            debug(stmt);
            gimple_set_visited(stmt, false);
        }
        printf("====================================\n");

        FOR_EACH_BB_FN(bb, fun)
        for (gsi = gsi_start_bb(bb); !gsi_end_p(gsi); gsi_next(&gsi)) {
            gimple* stmt = gsi_stmt(gsi);
            if (gimple_code(stmt) != GIMPLE_RETURN) {
                continue;
            }
            mark_stmt_visited(stmt);
        }

        FOR_EACH_BB_FN(bb, fun)
        for (gsi = gsi_start_bb(bb); !gsi_end_p(gsi);) {
            gimple* stmt = gsi_stmt(gsi);
            if (!gimple_visited_p(stmt)) {
                printf("remove: \n");
                debug(stmt);
                gsi_remove(&gsi, true);
            } else {
                gsi_next(&gsi);
            }
        }
        return 0;
    }

    virtual my_pass* clone() override { return this; }
};

struct register_pass_info my_passinfo {
    .pass = new my_pass(g), .reference_pass_name = "ssa",
    .ref_pass_instance_number = 1, .pos_op = PASS_POS_INSERT_AFTER
};

void register_callbacks(const char* base_name) {
    register_callback(base_name, PLUGIN_PASS_MANAGER_SETUP, NULL, &my_passinfo);
}
