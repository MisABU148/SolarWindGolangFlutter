package main

import (
	"backend/controller"
	"backend/db"
	"backend/repository"
	"backend/service"
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	dbConn, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	if err := db.InitSchema(dbConn); err != nil {
		log.Fatal("Schema init error:", err)
	}

	if err := db.SeedInitialData(dbConn, "./db/json/cities-names.json", "./db/json/sports.json"); err != nil {
		log.Fatal("Seeding error:", err)
	}

	userRepo := &repository.UserRepository{DB: dbConn}
	userService := &service.UserService{Repo: userRepo}
	userController := &controller.UserController{Service: userService}

	cityRepo := &repository.CityRepository{DB: dbConn}
	cityService := &service.CityService{Repo: cityRepo}
	cityController := &controller.CityController{Service: cityService}

	sportRepo := &repository.SportRepository{DB: dbConn}
	sportService := &service.SportService{Repo: sportRepo}
	sportController := &controller.SportController{Service: sportService}

	http.HandleFunc("/api/me", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userController.GetUserHandler(w, r)
		case http.MethodDelete:
			userController.DeleteUserHandler(w, r)
		case http.MethodPut:
			userController.UpdateUserHandler(w, r)
		case http.MethodPost:
			userController.CreateUserHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/getUsers", userController.GetAllUsersHandler)

	http.HandleFunc("/api/cities", cityController.GetCitiesHandler)
	http.HandleFunc("/api/cities/pagination", cityController.GetPaginatedCitiesHandler)
	http.HandleFunc("/api/cities/search", cityController.SearchCitiesHandler)

	http.HandleFunc("/api/sports", sportController.GetSportsHandler)
	http.HandleFunc("/api/sports/pagination", sportController.GetPaginatedSportsHandler)
	http.HandleFunc("/api/sports/search", sportController.SearchSportsHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
