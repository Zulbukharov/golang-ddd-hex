package adding

import (
	"fmt"
	"log"
)

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// Done means finished processing successfully
	Done Event = iota
	// PostAlreadyExist means the given beer is a duplicate of an existing one
	PostAlreadyExist
	// Failed means processing did not finish successfully
	Failed
)

// ErrDuplicate defines the error message
var ErrDuplicate = fmt.Errorf("Post already exist")

// GetMeaning returns string represenation of event
func (e Event) GetMeaning() string {
	var mean string

	mean = "Undefined"
	switch e {
	case Done:
		mean = "Done"
	case PostAlreadyExist:
		mean = "Post Already Exist"
	case Failed:
		mean = "Failed"
	}
	return mean
}

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
	log.Printf("Add post service\n")
	return s.tR.AddPost(u)
}
