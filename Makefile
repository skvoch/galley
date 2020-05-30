.PH0NY: build

build:
	go build -v ./cmd/main
	go build -v ./cmd/logger

test:
	go test -v -race ./...

.DEFAULT_GOAL := build
