package controller

import (
	"backend/model"
	"backend/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserController struct {
	Service *service.UserService
}

func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var dto model.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := uc.Service.CreateUser(dto)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}

func (uc *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := uc.Service.GetByUserID(id)
	// fmt.Print(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("User with ID %d not found", id), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = uc.Service.DeleteUserByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("User with ID %d not found", id), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var dto model.UserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := uc.Service.UpdateUser(dto); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uc.Service.GetUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (uc *UserController) CreateDeckHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем id из query-параметра
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Получаем пользователей из сервиса
	users, err := uc.Service.CreateDeckForUser(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create deck: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
