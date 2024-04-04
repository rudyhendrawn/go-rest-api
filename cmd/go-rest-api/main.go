package main

import (
	"go-rest-api/internal/app/handler"
	"go-rest-api/internal/app/middleware"
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

	// Use the middleware
	r.Use(middleware.Logging)
	r.Use(middleware.ErrorHandler)
	r.Use(middleware.CORSMiddleware)

	// Define the routes that using Error Handler
	// r.Handle("/users", middleware.ErrorHandlerFunc(userHandler.GetAllUsers)).Methods("GET")
	// r.Handle("/users/{id}", middleware.ErrorHandlerFunc(userHandler.GetUserByID)).Methods("GET")
	// r.Handle("/users", middleware.ErrorHandlerFunc(userHandler.CreateUser)).Methods("POST")
	// r.Handle("/posts", middleware.ErrorHandlerFunc(postHandler.GetAllPosts)).Methods("GET")
	// r.Handle("/posts/{id}", middleware.ErrorHandlerFunc(postHandler.GetPostByID)).Methods("GET")
	// r.Handle("/posts", middleware.ErrorHandlerFunc(postHandler.CreatePost)).Methods("POST")

	r.Handle("/createuser", middleware.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := userHandler.CreateUser(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))).Methods("POST")

	r.Handle("/user/{id}", middleware.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := userHandler.GetUserByID(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))).Methods("GET")

	r.Handle("/users", middleware.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := userHandler.GetAllUsers(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))).Methods("GET")

	r.Handle("/createpost", middleware.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := postHandler.CreatePost(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))).Methods("POST")

	r.Handle("/post/{id}", middleware.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := postHandler.GetPostByID(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))).Methods("GET")

	r.Handle("/posts", middleware.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := postHandler.GetAllPosts(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))).Methods("GET")

	// Start the server, listening on some port define in variable port
	var port = 1122
	log.Println("Starting server on: ", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	// defer db.Close()
}

// What to do next?
// 1. Middleware
//  - Logging (v)
//  - Error handling (v)
//  - CORS
//  - Authentication
//  - Authorization
