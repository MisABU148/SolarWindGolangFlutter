package controller

import (
	"backend/model"
	"backend/service"
	"encoding/json"
	"net/http"
)

type LikesController struct {
	Service *service.LikesService
}

func NewLikesController(s *service.LikesService) *LikesController {
	return &LikesController{Service: s}
}

func (lc *LikesController) HandleLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var dto model.LikeDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if dto.LikerID == 0 || dto.LikedID == 0 {
		http.Error(w, "liker_id and liked_id must be positive", http.StatusBadRequest)
		return
	}

	err = lc.Service.SaveOrUpdateDecision(r.Context(), &dto)
	if err != nil {
		http.Error(w, "Failed to save like decision: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Like decision saved"}`))
}
