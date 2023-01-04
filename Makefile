# BUILD=$(shell git rev-list -1 HEAD)

run:
	go run .\cmd\api\main.go

clean:
	go mod tidy

build-api:
	go mod tidy
	go mod download
	go build ./cmd/api/main.go

build-site:
	go mod tidy
	go mod download
	go build ./cmd/web/main.go

# build-sync:
# 	go mod tidy
# 	go mod download
# 	go build ./cmd/api/sync.go

docker:
	docker compose push