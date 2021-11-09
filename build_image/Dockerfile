FROM debian:latest

### Dependencies versions
ARG golang_version=1.17.2

##### Repos
RUN  apt-get update

##### General deps and build tools
RUN apt-get install --no-install-recommends -y \
    clang \
    lld \
    git-core \
    python3 \
    openjdk-11-jre-headless \
    make \
    wget \
    tar \
    zip \
    curl \
    less \
    nano \
    gcc

#### Installing GO
RUN url="https://golang.org/dl/go${golang_version}.linux-amd64.tar.gz"; goRelSha256='f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676' ; \
 	wget -O go.tgz "$url"; 	echo "${goRelSha256} *go.tgz" | sha256sum -c -; 	tar -C /usr/local -xzf go.tgz; 	rm go.tgz; \
    export PATH="/usr/local/go/bin:$PATH"; 	go version

ENV GOPATH=/go
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
ENV PATH=$PATH:/go/bin:/usr/local/go/bin


##### Setup clang as a compiler
# This is not really much necessary as clang is the only compiler installed,
# but let's do just in case
RUN update-alternatives --install /usr/bin/cc cc /usr/bin/clang 100 && \
    update-alternatives --install /usr/bin/c89 c89 /usr/bin/clang 100 && \
    update-alternatives --install /usr/bin/c99 c99 /usr/bin/clang 100 && \
    update-alternatives --install /usr/bin/c++ c++ /usr/bin/clang++ 100 && \
    update-alternatives --install /usr/bin/ld ld /usr/bin/lld 100

##### Install Bazel
RUN go get github.com/bazelbuild/bazelisk && \
    ln -s $(go env GOPATH)/bin/bazelisk /usr/bin/bazel
