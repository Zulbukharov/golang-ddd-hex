package userpolicy

// Service ...
type Service interface {
	CanAddPost(userID uint) bool
	//isAuthentificated(user User) bool
}

type Repository interface {
	CanAddPost(userID uint) bool
}

type service struct {
	authR Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{r}
}

// CanAddPost ...
func (s *service) CanAddPost(userID uint) bool {
	return s.authR.CanAddPost(userID)
}

//public MyBusinessProcess performSomeStuff(MyBusinessProcess input) {
//// We assume SecurityContext is a thread-local class that contains information
//// about the current user.
//if (!SecurityContext.isLoggedOn()) { //
//throw new AuthenticationException("No user logged on");
//}
//if (!SecurityContext.holdsRole("ROLE_BUSINESS_PROCESSOR")) { //
//throw new AccessDeniedException("Insufficient privileges");
//}
//
//var customer = customerRepository.findById(input.getCustomerId())
//.orElseThrow( () -> new CustomerNotFoundException(input.getCustomerId()));
//var someResult = myDomainService.performABusinessOperation(customer);
//customer = customerRepository.save(customer);
//return input.updateMyBusinessProcessWithResult(someResult);
//}
