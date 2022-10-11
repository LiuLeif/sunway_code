load("@bazel_tools//tools/cpp:cc_toolchain_config_lib.bzl", "tool_path")

def _impl(ctx):
    tool_paths = [ # NEW
        tool_path(
            name = "gcc",
            path = "/opt/riscv/bin/riscv64-unknown-linux-gnu-gcc",
        ),
        tool_path(
            name = "ld",
            path = "/opt/riscv/bin/riscv64-unknown-linux-gnu-ld",
        ),
        tool_path(
            name = "ar",
            path = "/opt/riscv/bin/riscv64-unknown-linux-gnu-ar",
        ),
        tool_path(
            name = "cpp",
            path = "/bin/false",
        ),
        tool_path(
            name = "gcov",
            path = "/bin/false",
        ),
        tool_path(
            name = "nm",
            path = "/bin/false",
        ),
        tool_path(
            name = "objdump",
            path = "/bin/false",
        ),
        tool_path(
            name = "strip",
            path = "/bin/false",
        ),
    ]
    return cc_common.create_cc_toolchain_config_info(
        ctx = ctx,
        cxx_builtin_include_directories = [
          "/opt/riscv/sysroot/usr/include",
          "/opt/riscv/riscv64-unknown-linux-gnu/include",
          "/opt/riscv/lib/gcc/riscv64-unknown-linux-gnu",
        ],
        toolchain_identifier = "local",
        host_system_name = "local",
        target_system_name = "local",
        target_cpu = "rv64g",
        target_libc = "unknown",
        compiler = "gcc",
        abi_version = "unknown",
        abi_libc_version = "unknown",
        tool_paths = tool_paths,
    )

cc_toolchain_config = rule(
    implementation = _impl,
    attrs = {},
    provides = [CcToolchainConfigInfo],
)
