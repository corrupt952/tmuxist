VERSION ?= 0.0.7
LDFLAGS ?= -ldflags "-s -w -X 'tmuxist/command.Version=$(VERSION)'"

# HACK: make [target] [ARGS...]
ARGS = $(filter-out $@,$(MAKECMDGOALS))

# HACK: nothing undefined target
%:
	@:

all: run

run:
	go run $(LDFLAGS) . $(ARGS)

fmt:
	@go fmt ./...

test:
	@go test -v ./...

lint:
	@go list | xargs golint

clean:
	@rm -rf pkg/*
