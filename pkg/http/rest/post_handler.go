package rest

import (
	"encoding/json"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/deleting"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/auth"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/userpolicy"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/listing"
)

// PostHandler provides access to Post api methods.
type PostHandler interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
	EditPost(w http.ResponseWriter, r *http.Request)
	// FindPostByID()
	// UpdatePost()
	// DeletePost()
}

type postHandler struct {
	l listing.Service
	a adding.Service
	d deleting.Service
	u userpolicy.Service
	// logger
}

// NewPostHandler post handler
func NewPostHandler(l listing.Service, a adding.Service, d deleting.Service, u userpolicy.Service) PostHandler {
	return &postHandler{
		l: l,
		a: a,
		d: d,
		u: u,
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

func (h postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil || postID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	credentials := r.Context().Value("credentials").(*auth.AppClaims)
	userID := credentials.ID

	if allowed := h.u.IsOwnerOfPost(userID, uint(postID)); allowed == false {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if err := h.d.DeletePost(uint(postID)); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("post deleted.")
}

func (h postHandler) EditPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_, ok := params["id"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
