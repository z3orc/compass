run:
	go run ./cmd/main.go

docker:
	docker build -t registry.gitlab.com/z3orc/dynamic-rpc .
	docker push registry.gitlab.com/z3orc/dynamic-rpc