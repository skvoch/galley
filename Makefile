.PH0NY: build

build:
	go build -v ./cmd/main
	go build -v ./cmd/logger

.DEFAULT_GOAL := build
