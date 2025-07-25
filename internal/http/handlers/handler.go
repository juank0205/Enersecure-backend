package handlers

import (
	"github.com/juank0205/Enersecure-backend/internal/core"
	"github.com/juank0205/Enersecure-backend/internal/db"
	"github.com/juank0205/Enersecure-backend/internal/service"
)

type Handler struct {
	User core.UserService
}

func NewHandler(database *db.DB) *Handler {
	return &Handler{User: service.NewUserService(database)}
}
