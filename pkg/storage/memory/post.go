package memory

// Post defines the properties of a Post to be listed
type Post struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}
