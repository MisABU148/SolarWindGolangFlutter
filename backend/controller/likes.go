package controller

import (
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

func (c *LikesController) HandleLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload struct {
		LikerID int64 `json:"likerId"`
		LikedID int64 `json:"likedId"`
		IsLiked bool  `json:"isFirstLikes"`
	}

	// Декодируем тело запроса
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Простая валидация
	if payload.LikerID <= 0 || payload.LikedID <= 0 {
		http.Error(w, "Invalid user IDs", http.StatusBadRequest)
		return
	}

	// Логика обработки лайка
	if err := c.Service.SaveOrUpdateLike(payload.LikerID, payload.LikedID, payload.IsLiked); err != nil {
		http.Error(w, "Error processing like: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
