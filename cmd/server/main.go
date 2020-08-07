package main

import (
	"fmt"
	auth2 "github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/auth"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest/middleware"
	"log"
	"net/http"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/listing"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/login"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/register"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/storage/postgres"
)

func main() {
	c, err := postgres.NewStorage("localhost", "5432", "adm", "1234", "blog")
	if err != nil {
		log.Printf("%v", err)
		return
	}
	postRepo := postgres.NewPostRepository(c)
	userRepo := postgres.NewUserRepository(c)

	adder := adding.NewService(postRepo)
	listing := listing.NewService(postRepo)
	login := login.NewService(userRepo)
	register := register.NewService(userRepo)

	auth := auth2.NewAuthenticator("ok")
	middleware := middleware.NewRules(auth)

	postHandler := rest.NewPostHandler(listing, adder)
	userHandler := rest.NewUserHandler(login, register, auth)

	server := &http.Server{
		Addr:    fmt.Sprint(":8000"),
		Handler: rest.Route(postHandler, userHandler, middleware),
	}
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed ! ")
	}
}
