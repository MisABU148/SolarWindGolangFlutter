package controller

import (
	"backend/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type AuthController struct {
	UserService *service.UserService
	AuthCache   *service.AuthCache
}

func NewAuthController(
	userService *service.UserService,
	authCache *service.AuthCache,
) *AuthController {
	return &AuthController{
		UserService: userService,
		AuthCache:   authCache,
	}
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

	a.AuthCache.SetCode(code, tgID)
	w.WriteHeader(http.StatusOK)
}

func (a *AuthController) GetTokenByCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	codeStr := r.URL.Query().Get("code")
	code, err := strconv.ParseInt(codeStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid code", http.StatusBadRequest)
		return
	}

	tgID, ok := a.AuthCache.GetUserIDByCode(code)
	if !ok {
		http.Error(w, "Code expired or invalid", http.StatusUnauthorized)
		return
	}

	result, err := a.UserService.CreateOrGetTelegramUser(tgID)
	if err != nil {
		http.Error(w, "User operation failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := service.GenerateJWT(result.UserID)
	if err != nil {
		http.Error(w, "JWT generation failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":   token,
		"user_id": result.UserID,
		"is_new":  result.IsNewUser,
	})
}
