package register

// Service ...
type Service interface {
	Register(User) error
}

// Repository ...
type Repository interface {
	Register(User) error
}

type service struct {
	lR Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{r}
}

// Login ...
func (s *service) Register(u User) error {
	// input validation
	return s.lR.Register(u)
}
