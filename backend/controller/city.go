// controller/city.go
package controller

import (
	"backend/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type CityController struct {
	Service *service.CityService
}

func (cc *CityController) GetCitiesHandler(w http.ResponseWriter, r *http.Request) {
	cities, err := cc.Service.GetCities()
	if err != nil {
		http.Error(w, "Failed to fetch cities", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cities)
}

func (cc *CityController) GetPaginatedCitiesHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	cities, err := cc.Service.GetPaginatedCities(page, size)
	if err != nil {
		http.Error(w, "Failed to fetch paginated cities", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cities)
}

func (cc *CityController) SearchCitiesHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing search parameter", http.StatusBadRequest)
		return
	}
	cities, err := cc.Service.SearchCities(word)
	if err != nil {
		http.Error(w, "Failed to search cities", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cities)
}
