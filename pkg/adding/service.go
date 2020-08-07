package adding

import (
	"fmt"
)

// ErrDuplicate defines the error message
var ErrDuplicate = fmt.Errorf("Post already exist")

// Service provides Post adding operations.
type Service interface {
	AddPost(Post) error
}

// Repository provides access to Post repository.
type Repository interface {
	AddPost(Post) error
}

type service struct {
	tR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddPost adds the given Post to the database
func (s *service) AddPost(u Post) error {
	// any validation can be done here
	return s.tR.AddPost(u)
}
