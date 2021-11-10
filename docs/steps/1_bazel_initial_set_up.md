# Initial Project Set Up

In this instruction you'll get information how to set up bazel files for initial build of go code.

# Specify Bazel Version

First need to specify bazel version been used to build project. For that need to create `.bazelversion` file at the root
directory of this project. `.bazelversion` file will be used by `bazelisk` to install proper version of `bazel`, read
more about this
at [bazelisk documentation](https://github.com/bazelbuild/bazelisk#how-does-bazelisk-know-which-bazel-version-to-run).

Example `.bazelversion` content:

```
5.0.0-pre.20210728.1
```

# Create Workspace

A workspace is a directory on your filesystem that contains the source files for the software you want to build, as well
as symbolic links to directories that contain the build outputs. Each workspace directory has a text file named
WORKSPACE which may be empty, or may contain references to external dependencies required to build the outputs.

You can read more about workspace
at [bazel documentation](https://docs.bazel.build/versions/4.2.1/build-ref.html#workspace)

To create `bazel` workspace - create file `WORKSPACE` at the root of the project with following content:

```
# Unique Workspace Name
workspace(name = "com_github_dgodyna_bazel_playground")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Bazel Rules For Go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2b1641428dff9018f9e85c0384f03ec6c10660d935b750e3fa1492a281a53b0f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
    ],
)

# Bazel Rules For gazelle
http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.17.2")

gazelle_dependencies()
```

We'll get created workspace with following rules:

* [rules_go](https://github.com/bazelbuild/rules_go) - rules which support go integration. Read its documentation for
  more info.
* [bazel-gazelle](https://github.com/bazelbuild/bazel-gazelle) - rules which will generate Build files for go packages.
  Read ots documentation for more info.

# Create Root Package

We have to create top level [bazel package](https://docs.bazel.build/versions/4.2.1/build-ref.html#packages). For that
file `BUILD.bazel` at application root directory with following content:

```
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/dgodyna/bazel-playground
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
```

This will generate new BUILD.bazel files for your project.

Now need to run `bazel` to generate all the project files. Read more about how `gazelle`
works [here](https://github.com/bazelbuild/bazel-gazelle).

```shell
bazel run //:gazelle
```

Now you have all the generated files, but need to update build files with go dependencies. For that use the following
command. It'll sync build targets dependencies with `go.mod` file.

```shell
bazel run //:gazelle-update-repos
```

It'll generate Build files for all the go packages. Now you should be able to run build. Try to build binary by the
following command:

```shell
bazel build //cmd/busybox
```

You should see following output

```
INFO: Analyzed target //cmd/busybox:busybox (9 packages loaded, 182 targets configured).
INFO: Found 1 target...
Target //cmd/busybox:busybox up-to-date:
  bazel-bin/cmd/busybox/busybox_/busybox
INFO: Elapsed time: 12.516s, Critical Path: 2.31s
INFO: 10 processes: 3 internal, 7 darwin-sandbox.
INFO: Build completed successfully, 10 total actions
```

Now let's try to run compiled program:

```
⇒  bazel-bin/cmd/busybox/busybox_/busybox
2:22PM INF Version: unknown
2:22PM INF Program uptime: '553.444207ms'
2:22PM INF Program uptime: '1.635086508s'
^C2:22PM INF received cancellation signal 'interrupt'
```

And we can run tests for whole project:

```
bazel test //... 
```