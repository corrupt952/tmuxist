FROM golang:1.27@sha256:e0174e51e81218523251d85d248a90d24c3d5e81543b4f07a5d66229397db190

RUN apt-get update && \
    apt-get install -y --no-install-recommends tmux && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV XDG_CACHE_HOME=/tmp/.cache
ENV GO111MODULE on

WORKDIR /go/src/app
RUN go get github.com/mitchellh/gox \
    github.com/ramya-rao-a/go-outline \
    github.com/nsf/gocode \
    github.com/acroca/go-symbols \
    github.com/fatih/gomodifytags \
    github.com/josharian/impl \
    github.com/haya14busa/goplay/cmd/goplay \
    github.com/go-delve/delve/cmd/dlv \
    golang.org/x/tools/gopls \
    golang.org/x/lint/golint
