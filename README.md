# bazel-playground

Bazel playground for Go application (as of now). The aim of this project is to check how to convert existing go project
previously built by `make` to `bazel` and share this experience.

The repository contains sample application with following:

1) Command line application [cmd/busybox](cmd/busybox)` which will print uptime until server stop. Required to test
   docker packaging;
2) Command line application [cmd/greeter](cmd/greeter) which will greet the user. Created to check how bazel manage
   dependencies.
3) Some public libraries under [pkg](pkg);
4) Some private libraries under [internal](internal);
5) Test which is using `testdata` resources under [internal/resource](internal/resource);
6) `busybox` Application will be packaged in docker image;
7) `busybox` will be packaged in zip achieve with additional script to start application.

# Project Targets

- [x] Build Go Binaries;
- [x] Build Go Libraries;
- [ ] Lint Go Code;
- [x] Test Go Code;
- [x] Package Go binaries in zip archive with specific structure;
- [x] Package binaries in docker image;
- [x] Create guides how to perform all this steps from scratch;

# Guides

Navigate to [docs](docs) page to see guides and external documentation for bazel migration. You can also check
out `clean` branch to make changes by yourself from scratch.