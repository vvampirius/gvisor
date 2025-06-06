load("//tools:defs.bzl", "go_library", "go_test")
load("//tools/go_generics:defs.bzl", "go_template_instance")

package(default_applicable_licenses = ["//:license"])

licenses(["notice"])

go_template_instance(
    name = "dir_refs",
    out = "dir_refs.go",
    package = "sys",
    prefix = "dir",
    template = "//pkg/refs:refs_template",
    types = {
        "T": "dir",
    },
)

go_library(
    name = "sys",
    srcs = [
        "dir_refs.go",
        "kcov.go",
        "pci.go",
        "save_restore.go",
        "sys.go",
    ],
    visibility = ["//pkg/sentry:internal"],
    deps = [
        "//pkg/abi/linux",
        "//pkg/atomicbitops",
        "//pkg/context",
        "//pkg/coverage",
        "//pkg/errors/linuxerr",
        "//pkg/fspath",
        "//pkg/fsutil",
        "//pkg/log",
        "//pkg/refs",
        "//pkg/sentry/arch",
        "//pkg/sentry/fsimpl/host",
        "//pkg/sentry/fsimpl/kernfs",
        "//pkg/sentry/inet",
        "//pkg/sentry/kernel",
        "//pkg/sentry/kernel/auth",
        "//pkg/sentry/memmap",
        "//pkg/sentry/vfs",
        "//pkg/usermem",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)

go_test(
    name = "sys_test",
    srcs = ["sys_test.go"],
    library = ":sys",
)

go_test(
    name = "sys_integration_test",
    srcs = ["sys_integration_test.go"],
    deps = [
        ":sys",
        "//pkg/abi/linux",
        "//pkg/sentry/fsimpl/testutil",
        "//pkg/sentry/kernel",
        "//pkg/sentry/kernel/auth",
        "//pkg/sentry/vfs",
        "@com_github_google_go_cmp//cmp:go_default_library",
    ],
)
