package userpolicy

// Service ...
type Service interface {
	IsOwnerOfPost(userID uint, postID uint) bool
	//SetAdminRole(u User)
	//SetAuthorRole(u User)
	//SetUserRole(u User)

	//CanDeletePost(u User, postID uint) bool
	//CanDeleteComment(User) bool
	//CanReact(User) bool
}

type Repository interface {
	IsOwnerOfPost(userID uint, postID uint) bool
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
func (s *service) IsOwnerOfPost(userID uint, postID uint) bool {
	if userID == 0 || postID == 0 {
		return false
	}
	return s.authR.IsOwnerOfPost(userID, postID)
}
