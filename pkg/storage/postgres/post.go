package postgres

import (
	"database/sql"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/listing"
	"log"
)

// Post defines the properties of a Post to be listed
type Post struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

// PostRepository keeps data in postgres db
type PostRepository struct {
	db *sql.DB
}

// NewPostRepository return a new repo
func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

// AddPost saves the given Post to the repository
func (s *PostRepository) AddPost(u adding.Post) error {
	log.Printf("Add post storage")
	_, err := s.db.Exec("INSERT INTO posts(content, author_id) VALUES($1, $2);", u.Content, u.AuthorID)
	if err != nil {
		return adding.ErrDuplicate
	}
	return nil
}

// GetAllPosts returns all Posts from the storage
func (s *PostRepository) GetAllPosts() ([]listing.Post, error) {
	posts := make([]listing.Post, 0)

	var (
		id       uint
		content  string
		authorID uint
	)

	rows, err := s.db.Query("SELECT id, content, author_id FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &content, &authorID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, listing.Post{ID: id, Content: content, AuthorID: authorID})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
