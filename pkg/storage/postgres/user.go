package postgres

import (
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/login"
	"log"
)

// User defines the properties of a User to be listed
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login checks the given User
func (s *Storage) Login(u login.User) error {
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
