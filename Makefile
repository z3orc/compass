BINARY=compass

run:
	go run ./cmd/main.go -debug -ratelimit 0

test:
	go test ./...

build:
	rm ./build/$(BINARY) & go build -v -buildvcs=true -o ./build/$(BINARY) ./cmd/main.go

.PHONY: build