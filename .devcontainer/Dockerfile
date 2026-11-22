FROM golang:1.27@sha256:205c7636d74e0ffbe3d007162df7d74f859f5a8c8322dfb9f907977686c126cf

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
