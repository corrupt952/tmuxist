VERSION ?= 0.0.5
LDFLAGS	?= "-X github.com/corrupt952/tmuxist/version.Version=${VERSION}"

all: build

build:
	gox -ldflags=${LDFLAGS} -output="pkg/{{.OS}}_{{.Arch}}/{{.Dir}}" -osarch="darwin/amd64 linux/amd64 linux/386"

package: build
	cd pkg \
		&& find * -type d | xargs -I{} tar -zcvf tmuxist_${VERSION}_{}.tar.gz {}/tmuxist \
		&& find * -type d | xargs -I{} rm -rf {}

release:
	@ghr -t ${GITHUB_TOKEN} -u ${CIRCLE_PROJECT_USERNAME} -r ${CIRCLE_PROJECT_REPONAME} -c ${CIRCLE_SHA1} -delete ${VERSION} pkg

run:
	go run *.go

fmt:
	go fmt ./...

test:
	go test -v ./...

lint:
	go list | xargs golint

clean:
	rm -rf pkg/*

vars:
	echo ${VERSION}
	echo ${LDFLAGS}

setup:
	go get -u github.com/mitchellh/gox
	go get -u golang.org/x/lint/golint
