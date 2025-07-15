package controller

import (
	"backend/model"
	"backend/repository"
	"backend/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type AuthController struct {
	UserRepo *repository.UserRepository
}

func (a *AuthController) PostCode(w http.ResponseWriter, r *http.Request) {
	codeStr := r.URL.Query().Get("code")
	tgIDStr := r.URL.Query().Get("userId")

	code, err1 := strconv.ParseInt(codeStr, 10, 64)
	tgID, err2 := strconv.ParseInt(tgIDStr, 10, 64)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid params", http.StatusBadRequest)
		return
	}

	service.SetCode(code, tgID)
	w.WriteHeader(http.StatusOK)
}

func (a *AuthController) GetTokenByCode(w http.ResponseWriter, r *http.Request) {
	codeStr := r.URL.Query().Get("code")
	code, err := strconv.ParseInt(codeStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid code", http.StatusBadRequest)
		return
	}

	tgID, ok := service.GetUserIDByCode(code)
	if !ok {
		http.Error(w, "Code expired or invalid", http.StatusUnauthorized)
		return
	}

	// Check or create user
	user, err := a.UserRepo.GetByTelegramID(tgID)
	if err != nil {
		userID, err := a.UserRepo.CreateTelegramUser(model.User{
			TelegramID: tgID,
			UserName:   "tg_" + strconv.FormatInt(tgID, 10),
			Alias:      "telegram_user",
		})
		if err != nil {
			http.Error(w, "Could not create user", http.StatusInternalServerError)
			return
		}
		tgID = userID
	} else {
		tgID = user.ID
	}

	// Generate token
	token, err := service.GenerateJWT(tgID)
	if err != nil {
		http.Error(w, "JWT error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
