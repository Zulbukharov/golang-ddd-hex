package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"
)

// Storage keeps data in postgres db
type Storage struct {
	db *sql.DB
}

// NewStorage returns a new Postgres storage
func NewStorage(host, port, user, password, dbName string) (*Storage, error) {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}

// AddPost saves the given Post to the repository
func (s *Storage) AddPost(u adding.Post) error {
	_, err := s.db.Exec("INSERT INTO posts(content) VALUES($1);", u.Content)
	if err != nil {
		return adding.ErrDuplicate
	}
	return nil
}

// GetAllPosts returns all Posts from the storage
func (s *Storage) GetAllPosts() ([]listing.Post, error) {
	var posts []listing.Post

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
