package core

import "github.com/juank0205/Enersecure-backend/internal/models"

type UserService interface {
	Login(email string, password string) (int64, error)
	Register(client *models.Client) (bool, error)
}
