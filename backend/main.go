package main

import (
	"backend/controller"
	"backend/db"
	"backend/middleware"
	"backend/notifier"
	"backend/repository"
	"backend/service"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

func main() {
	// 	connStr := fmt.Sprintf(
	// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 		os.Getenv("DB_HOST"),
	// 		os.Getenv("DB_PORT"),
	// 		os.Getenv("DB_USER"),
	// 		os.Getenv("DB_PASSWORD"),
	// 		os.Getenv("DB_NAME"),
	// 	)

	connStr := "host=localhost port=5434 user=postgres password=postgres dbname=userdb sslmode=disable"

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

	if _, err := os.Stat(citiesPath); os.IsNotExist(err) {
		log.Println("Файл не найден:", citiesPath)
	} else if _, err := os.Stat(sportsPath); os.IsNotExist(err) {
		log.Println("Файл не найден:", sportsPath)
	} else {
		if err := db.SeedInitialData(dbConn, citiesPath, sportsPath); err != nil {
			log.Fatal("Seeding error:", err)
		}
	}

	cityRepo := &repository.CityRepository{DB: dbConn}
	sportRepo := &repository.SportRepository{DB: dbConn}
	userRepo := &repository.UserRepository{DB: dbConn}
	likesRepo := &repository.LikesRepository{DB: dbConn}
	userService := &service.UserService{
		Repo:      userRepo,
		CityRepo:  cityRepo,
		SportRepo: sportRepo,
		LikesRepo: likesRepo,
	}
	userController := &controller.UserController{Service: userService}

	cityService := &service.CityService{Repo: cityRepo}
	cityController := &controller.CityController{Service: cityService}

	sportService := &service.SportService{Repo: sportRepo}
	sportController := &controller.SportController{Service: sportService}

	matchNotifier := &notifier.HttpMatchNotifier{}

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

	http.HandleFunc("/api/likes", middleware.JWTAuthMiddleware(likesController.HandleLike))
	http.HandleFunc("/notification", middleware.JWTAuthMiddleware(userController.GetMatches))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
