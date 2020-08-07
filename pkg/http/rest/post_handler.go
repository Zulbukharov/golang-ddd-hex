package rest

import (
	"encoding/json"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/auth"
	"net/http"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/listing"
)

// PostHandler provides access to Post api methods.
type PostHandler interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	// FindPostByID()
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
	w.Header().Set("Content-Type", "application/json")
	list, err := h.l.GetAllPosts()

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

	credentials := r.Context().Value("credentials").(*auth.AppClaims)
	post.AuthorID = credentials.ID

	if err := h.a.AddPost(post); err != nil {
		http.Error(w, "Failed to add post", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New post added.")
}
