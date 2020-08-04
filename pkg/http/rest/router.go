package rest

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// ServeHTTP used to push context to request
func (a apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("1 ServerHTTP\n")
	a.h.ServeHTTP(w, r.WithContext(a.ctx))
}

type apiHandler struct {
	ctx context.Context
	h   http.Handler
}

// middleware simple middleware to push value to the context
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("2 Middleware\n")
		ctx := context.WithValue(r.Context(), "another one", 666)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Route returns an http handler for the api.
func Route(h PostHandler, u UserHandler) http.Handler {
	ctx := context.WithValue(context.Background(), "ok", 3)

	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.Handle("/posts", &apiHandler{ctx, middleware(http.HandlerFunc(h.GetPosts))}).Methods("GET")
	api.HandleFunc("/post", h.AddPost).Methods("POST")

	api.HandleFunc("/login", u.Login).Methods("POST")
	api.HandleFunc("/register", u.Register).Methods("POST")

	http.Handle("/", router)
	return router
}
