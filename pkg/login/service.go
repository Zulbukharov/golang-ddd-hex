package login

// Service ...
type Service interface {
	Login(User) (uint, error)
}

// Repository ...
type Repository interface {
	Login(User) (uint, error)
}

type service struct {
	lR Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{r}
}

// Login ...
func (s *service) Login(u User) (uint, error) {
	// input validation
	return s.lR.Login(u)
}
