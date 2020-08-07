package rest

import (
	"encoding/json"
	"fmt"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/auth"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/login"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/register"
	"net/http"
	"time"
)

// UserHandler provides access to User api methods.
type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	l    login.Service
	r    register.Service
	auth auth.Authenticator
	// logger
}

// NewUserHandler login handler
func NewUserHandler(l login.Service, r register.Service, auth auth.Authenticator) UserHandler {
	return &userHandler{l, r, auth}
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

	var id uint
	id, err = h.l.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	claims, err := h.auth.GenerateToken(id)
	http.SetCookie(w, &http.Cookie{
		Name:    "credentials",
		Value:   claims,
		Expires: time.Now().Add(72 * time.Hour),
	})
	json.NewEncoder(w).Encode(fmt.Sprintf("user logged in %v", claims))
}

func (h userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user register.User

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.r.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	claims, err := h.auth.GenerateToken(id)
	http.SetCookie(w, &http.Cookie{
		Name:    "credentials",
		Value:   claims,
		Expires: time.Now().Add(72 * time.Hour),
	})
	json.NewEncoder(w).Encode(fmt.Sprintf("user registered %v", claims))
}
