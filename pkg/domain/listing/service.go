package listing

// Service provides Ticket listing operations.
type Service interface {
	GetAllTickets() ([]Ticket, error)
}

// Repository provides access to Ticket repository.
type Repository interface {
	// GetAllTickets returns all tickets saved in storage.
	GetAllTickets() ([]Ticket, error)
}

type service struct {
	tR Repository
}

// NewService creates an list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetAllTickets returns all Tickets from the storage
func (s *service) GetAllTickets() ([]Ticket, error) {
	return s.tR.GetAllTickets()
}
