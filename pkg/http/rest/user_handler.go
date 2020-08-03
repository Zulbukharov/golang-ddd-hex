package rest

import (
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/login"
	"log"
	"net/http"
)

// UserHandler provides access to User api methods.
type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	l login.Service
	// logger
}

// NewUserHandler login handler
func NewUserHandler(l login.Service) UserHandler {
	return &userHandler{l}
}

// Login handler for POST /api/login requests
func (h userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user login.User

	user.Username = "az"
	user.Password = "pass"

	err := h.l.Login(user)
	if err != nil {
		log.Printf("err %v", err)
		return
	}
	log.Printf("logged in")
}
