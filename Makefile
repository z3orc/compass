# BUILD=$(shell git rev-list -1 HEAD)

run:
	go run .\cmd\main.go

clean:
	go mod tidy

build:
	go mod tidy
	go mod download
	go build ./cmd/main.go

docker:
	docker build -t registry.gitlab.com/z3orc/dynamic-rpc .
	docker push registry.gitlab.com/z3orc/dynamic-rpc