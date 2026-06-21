package handlers

import (
	"cdn-api/internal/services"
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type CreateUserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	user, err := h.Service.CreateUser(req.Name, req.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	res := CreateUserResponse{
		ID:   "123",
		Name: user.Name,
		Age:  user.Age,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
