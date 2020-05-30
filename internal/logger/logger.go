package logger

import (
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/skvoch/galley/internal/galley/model"
	"github.com/skvoch/galley/internal/logger/galley_client"
	"github.com/skvoch/galley/internal/logger/key_scanner"
	"time"
)

func New(firstName string, secondName string) *Logger {
	return &Logger{
		client: galley_client.New(),

		user: &model.User{
			FirstName:  firstName,
			SecondName: secondName,
		},

		scanner: key_scanner.New(time.Second * 5),
	}
}

type Logger struct {
	client  *galley_client.Client
	user    *model.User
	scanner *key_scanner.Scanner
}

func (l *Logger) Run() error {
	if err := l.handshake(); err != nil {
		logrus.Error(err)
	}

	for {
		count := <-l.scanner.GetCountChannel()
		fmt.Println(count)
	}

	return nil
}

func (l *Logger) handshake() error {
	hash := hex.EncodeToString([]byte(l.user.SecondName + l.user.SecondName))

	l.user.Hash = hash

	return l.client.Handshake(l.user)
}
