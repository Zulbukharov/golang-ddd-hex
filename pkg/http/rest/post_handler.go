package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"
)

// PostHandler provides access to Post api methods.
type PostHandler interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	// GetPostByID()
	// UpdatePost()
	// DeletePost()
}

type handler struct {
	l listing.Service
	a adding.Service
	// logger
}

// NewPostHandler post handler
func NewPostHandler(l listing.Service, a adding.Service) PostHandler {
	return &handler{
		l: l,
		a: a,
	}
}

// GetPosts returns a handler for GET /api/posts requests
func (h handler) GetPosts(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list, err := h.l.GetAllPosts()

		if err != nil {
			http.Error(w, "Failed to get posts", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(list)
}

// AddPost returns a handler for POST /api/post requests
func (h handler) AddPost(w http.ResponseWriter, r *http.Request) {
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
