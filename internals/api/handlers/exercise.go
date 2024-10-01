package handlers

import (
	"encoding/json"
	"net/http"

	"xyz-task-2/internals/services/recommendation"
)

type ExerciseHandler struct {
	service *recommendation.Service
}

func NewExerciseHandler(service *recommendation.Service) *ExerciseHandler {
	return &ExerciseHandler{service: service}
}

func (h *ExerciseHandler) GenerateExercise(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	recommendation, err := h.service.GetTopErrors(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendation)
}
