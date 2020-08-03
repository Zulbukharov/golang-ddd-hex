package postgres

import (
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"
	"log"
)

// Post defines the properties of a Post to be listed
type Post struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

// AddPost saves the given Post to the repository
func (s *Storage) AddPost(u adding.Post) error {
	log.Printf("Add post storage")
	_, err := s.db.Exec("INSERT INTO posts(content) VALUES($1);", u.Content)
	if err != nil {
		return adding.ErrDuplicate
	}
	return nil
}

// GetAllPosts returns all Posts from the storage
func (s *Storage) GetAllPosts() ([]listing.Post, error) {
	log.Printf("Storage get all posts")
	posts := make([]listing.Post, 0)

	var (
		id      uint
		content string
	)

	rows, err := s.db.Query("SELECT id, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &content)
		if err != nil {
			return nil, err
		}

		posts = append(posts, listing.Post{ID: id, Content: content})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
