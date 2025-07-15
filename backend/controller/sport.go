package controller

import (
	"backend/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type SportController struct {
	Service *service.SportService
}

func (sc *SportController) GetSportsHandler(w http.ResponseWriter, r *http.Request) {
	sports, err := sc.Service.GetSports()
	if err != nil {
		http.Error(w, "Failed to fetch sports", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sports)
}

func (sc *SportController) GetPaginatedSportsHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	sports, err := sc.Service.GetPaginatedSports(page, size)
	if err != nil {
		http.Error(w, "Failed to fetch paginated sports", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sports)
}

func (sc *SportController) SearchSportsHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing search parameter", http.StatusBadRequest)
		return
	}
	sports, err := sc.Service.SearchSports(word)
	if err != nil {
		http.Error(w, "Failed to search sports", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sports)
}
