BINARY=compass
VERSION=3.0.0
BUILD=`git rev-parse --short HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64 arm64

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

run:
	go run main.go

clean:
	go mod tidy

build:
	go build ${LDFLAGS} -o ./bin/${BINARY}

build_all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o "./bin/$(BINARY)-$(GOOS)-$(GOARCH)")))