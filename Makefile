BINARY=compass

run:
	go run ./cmd/main.go -debug -ratelimit 0

test:
	go test ./...

build:
	rm ./bin/$(BINARY) & go build -v -buildvcs=true -o ./bin/$(BINARY) ./cmd/main.go

.PHONY: build