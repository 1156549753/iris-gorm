package services

import (
	"fmt"
	"free/do"
	_ "free/models"
)

type userService struct {
	// do *dao
}

type UserService interface {
	Sigunp() string
}

func NewUserService() UserService {
	return &userService{}
}

var adminUser = do.NewAdminUser()

func (u userService) Sigunp() string {
	fmt.Println("2:")
	user := adminUser.Find()
	fmt.Println(user)
	return "mindsd"
}
