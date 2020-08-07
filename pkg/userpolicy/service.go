package userpolicy

// Service ...
type Service interface {
	IsOwnerOfPost(u User, postID uint) bool
	//SetAdminRole(u User)
	//SetAuthorRole(u User)
	//SetUserRole(u User)

	//CanDeletePost(u User, postID uint) bool
	//CanDeleteComment(User) bool
	//CanReact(User) bool
}

type Repository interface {
	IsOwnerOfPost(u User, postID uint) bool
	//CanDeleteComment(User) bool
	//CanReact(User) bool
}

type service struct {
	authR Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{r}
}

// CanAddPost ...
func (s *service) IsOwnerOfPost(u User, postID uint) bool {
	return s.authR.IsOwnerOfPost(u, postID)
}
