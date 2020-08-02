package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/http/rest"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/storage/postgres"
)

func main() {
	s, err := postgres.NewStorage("localhost", "54320", "adm", "supersecret", "blog")
	if err != nil {
		log.Printf("%v", err)
		return
	}
	adder := adding.NewService(s)
	listing := listing.NewService(s)

	server := &http.Server{
		Addr:    fmt.Sprint(":8000"),
		Handler: rest.New(listing, adder),
	}
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed ! ")
	}
}
