OUTPUT_PATH	?= tmuxist
BUILD_FLAGS	?= -a -v -race

all: build

dep:
	dep ensure

build: dep
	gox -output="pkg/{{.Dir}}_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 linux/386"

run: dep
	go run *.go

fmt:
	gofmt -w *.go

test: dep
	go test .

clean:
	rm -f pkg/tmuxist_*

vars:
	echo ${OUTPUT_PATH}
	echo ${BUILD_FLAGS}
