package logger

import (
	"encoding/hex"
	"github.com/skvoch/galley/internal/galley/model"
	"github.com/skvoch/galley/internal/logger/galley_client"
)

func New(firstName string, secondName string) *Logger {
	return &Logger{
		client: galley_client.New(),

		user: &model.User{
			FirstName:  firstName,
			SecondName: secondName,
		},
	}
}

type Logger struct {
	client *galley_client.Client
	user   *model.User
}

func (l *Logger) Run() error {
	return l.handshake()
}

func (l *Logger) handshake() error {
	hash := hex.EncodeToString([]byte(l.user.SecondName + l.user.SecondName))

	l.user.Hash = hash

	return l.client.Handshake(l.user)
}
