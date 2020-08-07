package postgres

import (
	"database/sql"
)

type UserPolicyRepository struct {
	db *sql.DB
}

// NewUserRepository return a new repo
func NewUserPolicyRepository(db *sql.DB) *UserPolicyRepository {
	return &UserPolicyRepository{db}
}

func (s *UserPolicyRepository) IsOwnerOfPost(userID, postID uint) bool {
	stmt, err := s.db.Prepare("SELECT COUNT(1) FROM posts WHERE author_id = $1 and id = $2;")
	if err != nil {
		return false
	}
	defer stmt.Close()

	var id uint
	err = stmt.QueryRow(userID, postID).Scan(&id)
	if id == 0 || err != nil {
		return false
	}
	return true
}
