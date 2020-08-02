package memory

import (
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/adding"
	"github.com/Zulbukharov/golang-ddd-hex/pkg/domain/listing"
)

// Storage Memory keeps data in memory
type Storage struct {
	Tickets []Ticket
}

// AddTicket saves the given Ticket to the repository
func (s *Storage) AddTicket(u adding.Ticket) error {

	for _, e := range s.Tickets {
		if e.Content == u.Content {
			return adding.ErrDuplicate
		}
	}

	newU := Ticket{
		ID:      uint(len(s.Tickets) + 1),
		Content: u.Content,
	}
	s.Tickets = append(s.Tickets, newU)
	return nil
}

// GetAllTickets returns all Tickets from the storage
func (s *Storage) GetAllTickets() ([]listing.Ticket, error) {
	var tickets []listing.Ticket

	for i := range s.Tickets {
		ticket := listing.Ticket{
			ID:      s.Tickets[i].ID,
			Content: s.Tickets[i].Content,
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}
