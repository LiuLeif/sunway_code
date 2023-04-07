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
#include "insn-modes.h"
#include <gimple.h>
#include <gimple-iterator.h>
#include <gimple-pretty-print.h>
#include <context.h>
#include <ssa.h>
#include <rtl.h>
// clang-format on

const pass_data my_pass_data = {
    .type = RTL_PASS,
    .name = "my_pass",
    .optinfo_flags = OPTGROUP_NONE,
    .tv_id = TV_NONE,
    .properties_required = PROP_gimple_any,
    .properties_provided = 0,
    .properties_destroyed = 0,
    .todo_flags_start = 0,
    .todo_flags_finish = 0,
};

struct my_pass : rtl_opt_pass {
   public:
    my_pass(gcc::context* ctxt) : rtl_opt_pass(my_pass_data, ctxt) {}

    virtual unsigned int execute(function* fun) override {
        if (strcmp(function_name(fun), "trace") == 0) {
            return 0;
        }
        rtx call = gen_rtx_CALL(
            VOIDmode,
            gen_rtx_MEM(FUNCTION_MODE, gen_rtx_SYMBOL_REF(Pmode, "trace")),
            const0_rtx);

        basic_block bb;
        rtx_insn* insn;
        FOR_EACH_BB_FN(bb, fun) {
            FOR_BB_INSNS(bb, insn) {
                // if (GET_CODE(insn) == NOTE) {
                // continue;
                // }
                // rtx_insn* call_tace = emit_call_insn_before(call, insn);
                // rtx_insn* setx = emit_insn_before(
                // gen_rtx_SET(
                // gen_rtx_REG(DImode, 5),
                // gen_rtx_SYMBOL_REF(DImode, "__main")),
                // call_tace);
                emit_call_insn_before(call, insn);
                goto out;
            }
        }
    out:
        return 0;
    }

    virtual my_pass* clone() override { return this; }
};

struct register_pass_info my_passinfo {
    .pass = new my_pass(g), .reference_pass_name = "expand",
    .ref_pass_instance_number = 1, .pos_op = PASS_POS_INSERT_AFTER
};

void register_callbacks(const char* base_name) {
    register_callback(base_name, PLUGIN_PASS_MANAGER_SETUP, NULL, &my_passinfo);
}
