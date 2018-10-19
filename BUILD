load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/joe-mann/gazelle-fork-demo
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/joe-mann/gazelle-fork-demo",
    visibility = ["//visibility:private"],
    deps = ["@com_github_pelletier_go_toml//:go_default_library"],
)

go_binary(
    name = "gazelle-fork-demo",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
