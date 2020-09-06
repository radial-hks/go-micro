package service

import (
	"errors"
)

type IUserService interface {
	GetName(userID int) string
	DeleteUser(userID int) error
}

type UserService struct{}

func (*UserService) GetName(userID int) string {
	if userID == 101 {
		return "radial"
	}
	return "hah"
}

func (*UserService) DeleteUser(userID int) error {
	if userID == 101 {
		return errors.New("-777")
	}
	return nil
}
