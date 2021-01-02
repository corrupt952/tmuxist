FROM circleci/golang:1.15
ENV XDG_CACHE_HOME=/tmp/.cache
WORKDIR /go/src/app
RUN go get -u github.com/mitchellh/gox \
    | go get -u golang.org/x/lint/golint
