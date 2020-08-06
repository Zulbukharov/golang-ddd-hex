package adding

// Post basic post struct
type Post struct {
	AuthorID uint   `json:"author_id"`
	Content  string `json:"content"`
}
