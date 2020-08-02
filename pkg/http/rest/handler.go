package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"
	"github.com/gorilla/mux"
)

// New returns an http handler for the api.
func New(l listing.Service, a adding.Service) http.Handler {
	router := mux.NewRouter()
	post := router.PathPrefix("/api").Subrouter()
	post.HandleFunc("/posts", getPosts(l)).Methods("GET")
	post.HandleFunc("/post", addPost(a)).Methods("POST")
	http.Handle("/", router)
	return router
}

// getPosts returns a handler for GET /api/posts requests
func getPosts(s listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list, _ := s.GetAllPosts()
		json.NewEncoder(w).Encode(list)
	}
}

// addPost returns a handler for POST /api/post requests
func addPost(a adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var post adding.Post

		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&post); err != nil {
			http.Error(w, "Failed to parse post", http.StatusBadRequest)
			return
		}

		if err := a.AddPost(post); err != nil {
			http.Error(w, "Failed to add post", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New post added.")
	}
}
