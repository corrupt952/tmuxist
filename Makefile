VERSION ?= 0.0.3
LDFLAGS	?= "-X main.version=${VERSION}"

all: build

dep:
	dep ensure

build: dep
	gox -ldflags=${LDFLAGS} -output="pkg/{{.OS}}_{{.Arch}}/{{.Dir}}" -osarch="darwin/amd64 linux/amd64 linux/386"

package: build
	cd pkg \
		&& find * -type d | xargs -I{} tar -zcvf tmuxist_${VERSION}_{}.tar.gz {}/tmuxist \
		&& find * -type d | xargs -I{} rm -rf {}

release:
	ghr -t ${GITHUB_TOKEN} -u ${CIRCLE_PROJECT_USERNAME} -r ${CIRCLE_PROJECT_REPONAME} -c ${CIRCLE_SHA1} -delete ${VERSION} pkg


run: dep
	go run *.go

fmt:
	gofmt -w *.go

test: dep
	go test .

clean:
	rm -rf pkg/*

vars:
	echo ${VERSION}
	echo ${LDFLAGS}
