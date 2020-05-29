.PH0NY: build

build:
	go build -v ./cmd/main
test:
	go test -v -race ./...

.DEFAULT_GOAL := build
