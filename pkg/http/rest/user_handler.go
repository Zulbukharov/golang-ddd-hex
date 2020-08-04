package rest

import (
	"encoding/json"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/login"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/register"
	"net/http"
)

// UserHandler provides access to User api methods.
type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	l login.Service
	r register.Service
	// logger
}

// NewUserHandler login handler
func NewUserHandler(l login.Service, r register.Service) UserHandler {
	return &userHandler{l, r}
}

// Login handler for POST /api/login requests
func (h userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user login.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.l.Login(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// some cookie
	json.NewEncoder(w).Encode("User logged in")
}

func (h userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user register.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.r.Register(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("New user registered.")
}
