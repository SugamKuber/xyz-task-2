.PHONY: build run test docker-build docker-run fmt

build:
	go build -o bin/server cmd/init.go cmd/main.go

run: build
	./bin/server

test:
	go test ./...

docker-build:
	docker build -t exercise-recommendation-system .

docker-run:
	docker-compose up --build

docker-stop:
	docker-compose down

fmt:
	@echo "Formatting Go code..."
	go fmt ./...

.DEFAULT_GOAL := run