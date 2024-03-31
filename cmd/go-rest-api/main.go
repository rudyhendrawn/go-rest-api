package main

import (
	"go-rest-api/internal/app/handler"
	"go-rest-api/internal/app/repository"
	"go-rest-api/internal/app/service"
	"go-rest-api/internal/pkg/config"
	"go-rest-api/internal/pkg/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// Load the application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading application configuration: %v", err)
	}

	// Connect to the database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Initialize the repository
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

	// Initialize the service
	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)

	// Initialize the handlers
	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)

	// Set up the router
	r := mux.NewRouter()
	r.HandleFunc("/createuser", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.GetUserByID).Methods("GET")
	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/createpost", postHandler.CreatePost).Methods("POST")
	r.HandleFunc("/post/{id}", postHandler.GetPostByID).Methods("GET")

	// Start the server, listening on some port define in variable port

	var port = 1122
	log.Println("Starting server on: ", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	// defer db.Close()
}
