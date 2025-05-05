package service

import (
	"errors"
	"todo_app/mock"
	"todo_app/model"
)

func Authenticate(username, password string) (*model.User, error) {
	for _, user := range mock.Users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, errors.New("kullanıcı adı veya şifre yanlış")
}
