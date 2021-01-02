UID = $(shell id -u)
GID = $(shell id -g)
export UID
export GID

VERSION ?= 0.0.7
LDFLAGS	?= "-X github.com/corrupt952/tmuxist/version.Version=$(VERSION)"

all: build

build: clean
	@docker-compose run --rm app \
		bash -c 'gox -ldflags=$(LDFLAGS) -output="pkg/{{.OS}}_{{.Arch}}/{{.Dir}}" -osarch="darwin/amd64 linux/amd64 linux/386"'

fmt:
	@docker-compose run --rm app \
		bash -c 'go fmt ./...'

test:
	@docker-compose run --rm app \
		bash -c 'go test -v ./...'

lint:
	@docker-compose run --rm app \
		bash -c 'go list | xargs golint'

clean:
	@rm -rf pkg/*

vars:
	@echo "VERSION: $(VERSION)"
	@echo "LDFLAGS: $(LDFLAGS)"
	@echo "UID: $(UID)"
	@echo "GID: $(GID)"

###
# for CI
ci_test:
	go test -v ./...

ci_build:
	gox -ldflags=$(LDFLAGS) -output="pkg/{{.OS}}_{{.Arch}}/{{.Dir}}" -osarch="darwin/amd64 linux/amd64 linux/386"

ci_package: ci_build
	cd pkg \
		&& find * -type d | xargs -I{} tar -zcvf tmuxist_$(VERSION)_{}.tar.gz {}/tmuxist \
		&& find * -type d | xargs -I{} rm -rf {}

ci_release:
	ghr -t ${GITHUB_TOKEN} -u ${CIRCLE_PROJECT_USERNAME} -r ${CIRCLE_PROJECT_REPONAME} -c ${CIRCLE_SHA1} -delete ${VERSION} pkg
