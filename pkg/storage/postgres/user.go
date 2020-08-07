package postgres

import (
	"database/sql"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/login"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/register"
)

// User defines the properties of a User to be listed
type User struct {
	ID       uint
	Username string
	Password string
	RoleID   uint
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
func (s *UserRepository) Login(u login.User) (uint, error) {
	stmt, err := s.db.Prepare("SELECT id, username, password FROM users WHERE username = $1 and password = $2")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var (
		id       uint
		username string
		password string
	)
	err = stmt.QueryRow(u.Username, u.Password).Scan(&id, &username, &password)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (s *UserRepository) Register(u register.User) (uint, error) {
	stmt, err := s.db.Prepare("INSERT INTO users(username, password, role_id) VALUES ($1, $2, 1) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var id uint
	err = stmt.QueryRow(u.Username, u.Password).Scan(&id)
	return id, err
}
