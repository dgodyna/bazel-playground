# Docker Packaging

In this step we'll create docker image with generated application.

# Add Docker Packaging Rules

To support docker packaging following [rules](https://github.com/bazelbuild/rules_docker) should be added. It'll allow
building, pulling and pushing Docker/OCI images.

To add this repository rules add following to `WORKSPACE`:

```
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "92779d3445e7bdc79b961030b996cb0c91820ade7ffa7edca69273f404b085d5",
    strip_prefix = "rules_docker-0.20.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.20.0/rules_docker-v0.20.0.tar.gz"],
)


load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

```

Now let's load base image which will be used as a baseline for application. Add following to `WORKSPACE`:

```
container_pull(
  name = "busy_box_base",
  registry = "gcr.io",
  repository = "distroless/base-debian11",
  digest = "sha256:0530d193888bcd7bd0376c8b34178ea03ddb0b2b18caf265135b6d3a393c8d05",
)
```

# Add Packaging

Let's add steps to build our image. For this add following to root `BUILD.bazel`:

```
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")

# Name of docker registry to push image
REGISTRY = "kubernetes.docker.internal:5001"

# Build docker image with busybox application
container_image(
    name = "busy_box_image",
    base = "@busy_box_base//image",
    entrypoint = "busybox",
    files = ["//cmd/busybox"],
)

# Push it into specified registry
container_push(
    name = "busy_box_image_publish",
    format = "Docker",
    image = ":busy_box_image",
    registry = REGISTRY,
    repository = "bazel-playground",
    tag = VERSION,
)

```

Now let's run this steps:

```
⇒  bazelisk build //:busy_box_image_publish
INFO: Analyzed target //:busy_box_image_publish (77 packages loaded, 516 targets configured).
INFO: Found 1 target...
Target //:busy_box_image_publish up-to-date:
  bazel-bin/busy_box_image_publish.digest
  bazel-bin/busy_box_image_publish
INFO: Elapsed time: 3.619s, Critical Path: 2.45s
INFO: 14 processes: 1 internal, 13 darwin-sandbox.
INFO: Build completed successfully, 14 total actions
```

Now we have script to push image `bazel-bin/busy_box_image_publish` which can be executed to push image into specified
registry.

