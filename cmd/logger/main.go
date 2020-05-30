package main

import (
	logger2 "github.com/skvoch/galley/internal/logger"
	"time"
)

func main() {
	logger := logger2.New("Dmitriy", "Prokudin", time.Second*5)

	logger.Run()
}
