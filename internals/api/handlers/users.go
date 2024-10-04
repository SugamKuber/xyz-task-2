package handlers

import (
	"encoding/json"
	"net/http"

	"xyz-task-2/internals/services/users"
)

type UserHandler struct {
	service *users.Service
}

func NewUserHandler(service *users.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
