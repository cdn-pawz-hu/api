package handlers

import (
	"cdn-api/internal/dto"
	"cdn-api/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service  services.UserService
	validate *validator.Validate
}

func NewUserHandler(service services.UserService, v *validator.Validate) *UserHandler {
	return &UserHandler{service: service, validate: v}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.RegisterUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}
