# bazel-playground

Bazel playground for Go language. This project aims to check how to convert existing go project built by `make`
to `bazel` and collect some expertise.

The repository contains sample application with following:

1) Command line application [cmd/busybox](cmd/busybox)` which will print uptime until server stop. Required to test
   docker image creation;
2) Command line application [cmd/greeter](cmd/greeter) which will greet the user. Created to check dependencies.
3) Some public libraries under [pkg](pkg);
4) Some private libraries under [internal](internal);
5) Test which is using resource under [internal/resource](internal/resource);
6) Application will be packaged in docker image;
7) Application will be packaged in zip achieve with additional script to start application.

# Project Targets

- [x] Build Go Binaries;
- [x] Build Go Libraries;
- [ ] Lint Go Code;
- [x] Test Go Code;
- [x] Package Go binaries in zip archive with specific structure;
- [x] Package binaries in docker image;
- [x] Create guides to do all this steps from scratch;

# Guides

Navigate to [docs](docs) page to see guides for bazel migration. You can also check out `clean` branch to make changes
by yourself from scratch.