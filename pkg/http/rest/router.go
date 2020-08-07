package rest

import (
	"context"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/middleware"
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

// Route returns an http handler for the api.
func Route(h PostHandler, u UserHandler) http.Handler {
	ctx := context.WithValue(context.Background(), "ok", 3)

	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/posts", h.GetPosts).Methods("GET")
	api.Handle("/post", &apiHandler{ctx, middleware.LoggedIn(http.HandlerFunc(h.AddPost))}).Methods("POST")

	api.HandleFunc("/login", u.Login).Methods("POST")
	api.HandleFunc("/register", u.Register).Methods("POST")

	http.Handle("/", router)
	return router
}
