package main

import (
	"log"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"
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

	adder.AddPost(adding.Post{Content: "ok"})
	err = adder.AddPost(adding.Post{Content: "ok"})
	log.Printf("%v\n", err)
	l, _ := listing.GetAllPosts()
	log.Printf("%v\n", l)
}
