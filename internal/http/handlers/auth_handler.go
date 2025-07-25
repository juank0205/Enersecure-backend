package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juank0205/Enersecure-backend/internal/http/auth"
	"github.com/juank0205/Enersecure-backend/internal/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	// Decode JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	clientId, err := h.User.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}

	var response LoginResponse
	token, err := auth.GenerateJWT(int64(clientId))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	response.Token = token

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type RegisterRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	GovernmentID uint   `json:"government_id"`
	PhoneNumber  string `json:"phone_number"`
	Password     string `json:"password"`
	Email        string `json:"email"`
}

type RegisterResponse struct {
	Success bool `json:"success"`
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client := models.Client{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		GovernmentID: req.GovernmentID,
		PhoneNumber:  req.PhoneNumber,
		Password:     req.Password,
		Email:        req.Email,
	}

	success, err := h.User.Register(&client)
	if err != nil {
		fmt.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(RegisterResponse{Success: success})
		return
	}

	json.NewEncoder(w).Encode(RegisterResponse{Success: success})
}
