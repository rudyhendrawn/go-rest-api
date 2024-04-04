package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-rest-api/internal/app/model"
	"go-rest-api/internal/app/service"

	"github.com/gorilla/mux"
)

// PostHandler represents the HTTP handler for managing posts.
type PostHandler struct {
	postService service.PostService
}

// NewPostHandler creates a new instance of PostHandler.
func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

// CreatePost handles POST requests to create a new post.
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) error {
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	if err := h.postService.CreatePost(r.Context(), &post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)

	return nil
}

// GetPostByID handles GET requests to fetch a post by ID.
func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	post, err := h.postService.GetPostByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return err
	}

	json.NewEncoder(w).Encode(post)

	return nil
}

// GetAllPosts handles GET request to fetch all posts.
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) error {
	posts, err := h.postService.GetAllPosts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

	return nil
}
