package main

import (
	"log"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/storage/memory"
)

func main() {
	s := new(memory.Storage)
	adder := adding.NewService(s)
	listing := listing.NewService(s)

	adder.AddPost(adding.Post{Content: "ok"})
	err := adder.AddPost(adding.Post{Content: "ok"})
	log.Printf("%v\n", err)
	l, _ := listing.GetAllPosts()
	log.Printf("%v\n", l)
}
