package listing

// Service provides Post listing operations.
type Service interface {
	GetAllPosts() ([]Post, error)
	//GetUserPosts(id uint) ([]Post, error)
}

// Repository provides access to Post repository.
type Repository interface {
	// GetAllPosts returns all Posts saved in storage.
	GetAllPosts() ([]Post, error)
	//GetUserPosts(id uint) ([]Post, error)
}

type service struct {
	tR Repository
}

// NewService creates an list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetAllPosts returns all Posts from the storage
func (s *service) GetAllPosts() ([]Post, error) {
	return s.tR.GetAllPosts()
}

//func (s *service) GetUserPosts(id uint) ([]Post, error) {
//	return s.tR.GetUserPosts(id)
//}
