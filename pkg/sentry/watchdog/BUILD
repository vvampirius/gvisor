load("//tools:defs.bzl", "go_library")

package(
    default_applicable_licenses = ["//:license"],
    licenses = ["notice"],
)

go_library(
    name = "watchdog",
    srcs = ["watchdog.go"],
    visibility = ["//:sandbox"],
    deps = [
        "//pkg/log",
        "//pkg/metric",
        "//pkg/sentry/kernel",
        "//pkg/sentry/ktime",
        "//pkg/sync",
    ],
)
