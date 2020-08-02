package listing

// Ticket defines the properties of a ticket to be listed
type Ticket struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}
