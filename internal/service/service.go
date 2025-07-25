package service

import (
	"log"

	"github.com/juank0205/Enersecure-backend/internal/core"
	"github.com/juank0205/Enersecure-backend/internal/db"
	"github.com/juank0205/Enersecure-backend/internal/models"
)

type userService struct{
	db *db.DB
}

func NewUserService(database *db.DB) core.UserService {
	return &userService{db: database}
}

func (u *userService) Login(email string, password string) (int64, error) {
	clientId, err := u.db.GetEmailAndPassword(email, password)
	if err == nil {
		log.Printf("userService.Login: %v", err)
	}
	return clientId, err
}

func (u *userService) Register(client *models.Client) (bool, error) {
	available, err := u.db.CheckEmailAvailable(client.Email)
	if err != nil || !available  {
		return false, err
	}
	client_id, err := u.db.RegisterClient(client)
	if err != nil || client_id == 0 {
		return false, err
	}
	return true, nil
}
