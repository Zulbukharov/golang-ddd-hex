package postgres

import (
	"database/sql"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/login"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/register"
	"log"
)

// User defines the properties of a User to be listed
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserRepository keeps data in postgres db
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository returns a new repo
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

// Login checks the given User
func (s *UserRepository) Login(u login.User) error {
	stmt, err := s.db.Prepare("SELECT id, username, password FROM users WHERE username = $1 and password = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var (
		id       uint
		username string
		password string
	)
	err = stmt.QueryRow(u.Username, u.Password).Scan(&id, &username, &password)
	if err != nil {
		return err
	}
	log.Printf("id, username, password %d %s %s", id, username, password)
	return err
}

func (s *UserRepository) Register(u register.User) error {
	stmt, err := s.db.Prepare("INSERT INTO users(username, password) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Username, u.Password)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	log.Printf("last user id %d\n", lastID)
	return err
}
