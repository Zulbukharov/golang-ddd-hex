package register

import "fmt"

// Service ...
type Service interface {
	Register(User) (uint, error)
}

// Repository ...
type Repository interface {
	Register(User) (uint, error)
}

type service struct {
	lR Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{r}
}

// Login ...
func (s *service) Register(u User) (uint, error) {
	// input validation
	if u.Username == "" || u.Password == "" {
		return 0, fmt.Errorf("invalid input data")
	}
	return s.lR.Register(u)
}
