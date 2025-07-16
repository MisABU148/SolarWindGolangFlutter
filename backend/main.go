package main

import (
	"backend/controller"
	"backend/db"
	"backend/middleware"
	"backend/notifier"
	"backend/repository"
	"backend/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Printf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	if err := db.InitSchema(dbConn); err != nil {
		log.Fatal("Schema init error:", err)
	}

	citiesPath := "./db/json/cities-names.json"
	sportsPath := "./db/json/sports.json"

	// Проверяем, существуют ли файлы
	if _, err := os.Stat(citiesPath); os.IsNotExist(err) {
		log.Println("Файл не найден:", citiesPath)
	} else if _, err := os.Stat(sportsPath); os.IsNotExist(err) {
		log.Println("Файл не найден:", sportsPath)
	} else {
		if err := db.SeedInitialData(dbConn, citiesPath, sportsPath); err != nil {
			log.Fatal("Seeding error:", err)
		}
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

	likesRepo := &repository.LikesRepository{DB: dbConn}
	matchNotifier := &notifier.HttpMatchNotifier{} // твоя реализация интерфейса MatchNotifier

	likesService := &service.LikesService{
		Repo:     likesRepo,
		Notifier: matchNotifier,
	}
	likesController := controller.NewLikesController(likesService)

	authCache := service.NewAuthCache(userRepo)

	authController := controller.NewAuthController(userService, authCache)

	http.HandleFunc("/custom-auth", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			authController.GetTokenByCode(w, r)
		case http.MethodPost:
			authController.PostCode(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/me", middleware.JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
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
	}))

	http.HandleFunc("/api/getUsers", middleware.JWTAuthMiddleware(userController.GetAllUsersHandler))
	http.HandleFunc("/api/create-deck", middleware.JWTAuthMiddleware(userController.CreateDeckHandler))

	http.HandleFunc("/api/cities", cityController.GetCitiesHandler)
	http.HandleFunc("/api/cities/pagination", cityController.GetPaginatedCitiesHandler)
	http.HandleFunc("/api/cities/search", cityController.SearchCitiesHandler)

	http.HandleFunc("/api/sports", sportController.GetSportsHandler)
	http.HandleFunc("/api/sports/pagination", sportController.GetPaginatedSportsHandler)
	http.HandleFunc("/api/sports/search", sportController.SearchSportsHandler)

	http.HandleFunc("/api/likes", likesController.HandleLike)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
