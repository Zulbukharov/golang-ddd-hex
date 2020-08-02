package adding

import "fmt"

// Event defines possible outcomes from the "adding actor"
type Event int

const (
	// Done means finished processing successfully
	Done Event = iota
	// TicketAlreadyExist means the given beer is a duplicate of an existing one
	TicketAlreadyExist
	// Failed means processing did not finish successfully
	Failed
)

// ErrDuplicate defines the error message
var ErrDuplicate = fmt.Errorf("Ticket already exist")

// GetMeaning returns string represenation of event
func (e Event) GetMeaning() string {
	var mean string

	mean = "Undefined"
	switch e {
	case Done:
		mean = "Done"
	case TicketAlreadyExist:
		mean = "Ticket Already Exist"
	case Failed:
		mean = "Failed"
	}
	return mean
}

// Service provides Ticket adding operations.
type Service interface {
	AddTicket(Ticket) error
}

// Repository provides access to Ticket repository.
type Repository interface {
	AddTicket(Ticket) error
}

type service struct {
	tR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddTicket adds the given Ticket to the database
func (s *service) AddTicket(u Ticket) error {
	// any validation can be done here
	return s.tR.AddTicket(u)
}
