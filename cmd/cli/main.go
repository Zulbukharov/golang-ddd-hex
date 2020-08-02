package main

import (
	"github.com/Zulbukharov/golang-ddd-hex/pkg/storage/memory"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
)

func main() {
	s := new(memory.Storage)
	adder := adding.NewService(s)
	// listing := listing.NewService(s)

	adder.AddPost(adding.Post{Content: "sample1"})
	adder.AddPost(adding.Post{Content: "sample2"})
}
