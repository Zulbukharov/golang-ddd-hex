package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route returns an http handler for the api.
func Route(h PostHandler) http.Handler {
	router := mux.NewRouter()
	post := router.PathPrefix("/api").Subrouter()
	post.HandleFunc("/posts", h.GetPosts).Methods("GET")
	post.HandleFunc("/post", h.AddPost).Methods("POST")
	http.Handle("/", router)
	return router
}