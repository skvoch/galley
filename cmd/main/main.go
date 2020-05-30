package main

import (
	_ "github.com/bmizerany/pq"
	"github.com/sirupsen/logrus"
	"github.com/skvoch/galley/internal/galley/service"
)

func main() {
	s, err := service.New()

	if err != nil {
		logrus.Error(err)
	}

	s.Setup()
	s.Run()
}
