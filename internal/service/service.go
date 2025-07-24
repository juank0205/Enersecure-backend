package service

import (
	"github.com/juank0205/Enersecure-backend/internal/core"
)

type userService struct{}

func NewUserService() core.UserService {
	return &userService{}
}

func (s *userService) GetWelcomeMessage(name string) string {
	if name == "" {
		name = "Guest"
	}
	return "Hello, " + name + "!"
}
