package login

// Service ...
type Service interface {
	Login(User) error
}

// Repository ...
type Repository interface {
	Login(User) error
}

type service struct {
	lR Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{r}
}

// Login ...
func (s *service) Login(u User) error {
	// input validation
	return s.lR.Login(u)
}
