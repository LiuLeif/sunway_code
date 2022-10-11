cc_toolchain_suite(
    name = "toolchain",
    toolchains = {
        "rv64g": ":gnu_toolchain",
    },
)

filegroup(name = "empty")

cc_toolchain(
    name = "gnu_toolchain",
    toolchain_identifier = "gnu-toolchain",
    toolchain_config = ":gnu_toolchain_config",
    all_files = ":empty",
    compiler_files = ":empty",
    dwp_files = ":empty",
    linker_files = ":empty",
    objcopy_files = ":empty",
    strip_files = ":empty",
    supports_param_files = 0,
)

package(default_visibility = ["//visibility:public"])

load(":cc_toolchain_config.bzl", "cc_toolchain_config")

cc_toolchain_config(name = "gnu_toolchain_config")
