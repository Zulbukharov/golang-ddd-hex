package memory

import (
	"github.com/Zulbukharov/golang-ddd-hex/pkg/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/listing"
)

// Storage Memory keeps data in memory
type Storage struct {
	Posts []Post
}

// AddPost saves the given Post to the repository
func (s *Storage) AddPost(u adding.Post) error {

	for _, e := range s.Posts {
		if e.Content == u.Content {
			return adding.ErrDuplicate
		}
	}

	newU := Post{
		ID:      uint(len(s.Posts) + 1),
		Content: u.Content,
	}
	s.Posts = append(s.Posts, newU)
	return nil
}

// GetAllPosts returns all Posts from the storage
func (s *Storage) GetAllPosts() ([]listing.Post, error) {
	var posts []listing.Post

	for i := range s.Posts {
		Post := listing.Post{
			ID:      s.Posts[i].ID,
			Content: s.Posts[i].Content,
		}

		posts = append(posts, Post)
	}

	return posts, nil
}
