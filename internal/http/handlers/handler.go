package handlers

import (
	"net/http"

	"github.com/juank0205/Enersecure-backend/internal/core"
	"github.com/juank0205/Enersecure-backend/internal/db"
	"github.com/juank0205/Enersecure-backend/internal/service"
)

type Handler struct {
	User core.UserService
	DB *db.DB
}

func NewHandler(database *db.DB) *Handler {
	return &Handler{DB: database, User: service.NewUserService()}
}

func (h *Handler) Welcome(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	msg := h.User.GetWelcomeMessage(name)
	w.Write([]byte(msg))
}
