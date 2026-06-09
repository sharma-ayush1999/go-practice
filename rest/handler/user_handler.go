package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sharma-ayush1999/rest/usecase"
)

type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Name string `json:"name" validate:"required,min=2"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}

type UserHandler struct {
	createUser *usecase.CreateUserUseCase
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	output, err := h.createUser.Execute(r.Context(), usecase.CreateUserInput{
		Email: req.Email, Name: req.Name,
	})

	if err != nil {
		// Map use-case errors to HTTP status codes
		if strings.Contains(err.Error(), "validation") ||
		strings.Contains(err.Error(), "already"){
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateUserResponse{UserID: output.UserID})

}