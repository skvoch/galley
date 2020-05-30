package main

import (
	logger2 "github.com/skvoch/galley/internal/logger"
)

func main() {
	logger := logger2.New("Dmitriy", "Prokudin")

	logger.Run()
}
