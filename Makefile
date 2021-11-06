VERSION ?= 0.0.7
LDFLAGS ?= -ldflags "-s -w -X 'tmuxist/command.Version=$(VERSION)'"

ARGS = $(filter-out $@,$(MAKECMDGOALS))
%:
	@:

all: run

run:
	go run $(LDFLAGS) . $(ARGS)

fmt:
	go fmt ./...

test:
	go test -v ./...

lint:
	go list | xargs golint

clean:
	@rm -rf pkg/*
