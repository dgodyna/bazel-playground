# Set Up Environment

To operate with `bazel` we'll use [bazelisk](https://github.com/bazelbuild/bazelisk) which is easy to use wrapper
for `bazel`.

Now you have 2 ways - install everything locally, or process in container.

# Create Bazel Container

[Here](../../build_image/Dockerfile) you can find dockerfile which contains all that required for `bazel` build.

You can compile it from the following command executed from folder containing this DockerFile:

```shell
docker build . -t debian_bazel
```

Now you can use this container to compile your sources. For that just mount your source code to container directory.
Example:

```shell
docker run -i -t -v /Users/dgodyna/projects/bazel-playground:/bazel-playground debian_bazel
```

# Install Bazel Locally

To install `bazelisk` locally just follow steps
from [bazelisk repo](https://github.com/bazelbuild/bazelisk#installation).