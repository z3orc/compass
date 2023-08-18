BINARY=compass
VERSION=3.1.0
BUILD=`git rev-parse --short HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64 arm64

LDFLAGS=-ldflags "-X github.com/z3orc/dynamic-rpc/internal/env.Version=${VERSION} -X github.com/z3orc/dynamic-rpc/internal/env.Build=${BUILD}"

run:
	go run main.go

clean:
	go mod tidy

build:
	go build ${LDFLAGS} -o ./bin/${BINARY}

build-all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o "./bin/$(BINARY)-$(GOOS)-$(GOARCH)")))

build-docker:
	docker build --force-rm -t ${BINARY}:${VERSION}-${BUILD} -t ${BINARY}:latest . 