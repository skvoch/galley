package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	logger2 "github.com/skvoch/galley/internal/logger"
	"time"
)

func main() {

	firstName := flag.String("f", "", "Your first name")
	secondName := flag.String("s", "", "Your second name")
	flag.Parse()

	if len(*firstName) == 0 || len(*secondName) == 0 {
		fmt.Println("Use command line arguments for setting your name")
		fmt.Println(" -f \"Биба\" ")
		fmt.Println(" -s \"Бобович\" ")
		return
	}

	logrus.Info("Login as: ", *firstName, " ", *secondName)
	logger := logger2.New(*firstName, *secondName, time.Second*5)

	logger.Run()
}
