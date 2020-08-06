package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/listing"
)

// PostHandler provides access to Post api methods.
type PostHandler interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	// GetPostByID()
	// UpdatePost()
	// DeletePost()
}

type postHandler struct {
	l listing.Service
	a adding.Service
	// logger
}

// NewPostHandler post handler
func NewPostHandler(l listing.Service, a adding.Service) PostHandler {
	return &postHandler{
		l: l,
		a: a,
	}
}

// GetPosts handler for GET /api/posts requests
func (h postHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	z := r.Context().Value("ok")
	another := r.Context().Value("another one")
	log.Printf("%d Handler [from ServerHTTP context]\n", z)
	log.Printf("%d Handler [from middleware context]\n", another)
	w.Header().Set("Content-Type", "application/json")
	list, err := h.l.GetAllPosts()
	log.Printf("postHandler get posts")

	if err != nil {
		http.Error(w, "Failed to get posts", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(list)
}

// AddPost handler for POST /api/post requests
func (h postHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	var post adding.Post

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&post); err != nil {
		http.Error(w, "Failed to parse post", http.StatusBadRequest)
		return
	}

	log.Printf("Add post handler")
	if err := h.a.AddPost(post); err != nil {
		http.Error(w, "Failed to add post", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New post added.")
}
