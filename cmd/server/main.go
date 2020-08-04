package main

import (
	"fmt"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/login"
	"log"
	"net/http"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/register"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest"
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

	postHandler := rest.NewPostHandler(listing, adder)
	userHandler := rest.NewUserHandler(login, register)

	server := &http.Server{
		Addr:    fmt.Sprint(":8000"),
		Handler: rest.Route(postHandler, userHandler),
	}
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed ! ")
	}
}
