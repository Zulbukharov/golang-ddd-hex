package rest

import (
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

// Route returns an http handler for the api.
func Route(h PostHandler, u UserHandler, m middleware.Rules) http.Handler {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/posts", h.GetPosts).Methods("GET")
	api.Handle("/post", m.LoggedIn(http.HandlerFunc(h.AddPost))).Methods("POST")
	api.HandleFunc("/login", u.Login).Methods("POST")
	api.HandleFunc("/register", u.Register).Methods("POST")

	router.Use(accessControl)
	http.Handle("/", router)
	return router
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
