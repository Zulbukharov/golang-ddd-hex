package memory

// Ticket defines the properties of a Ticket to be listed
type Ticket struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}
