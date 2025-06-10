package user

import (
	"github.com/igorezka/auth/internal/service"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

// Implementation represents a user API implementation.
type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation crates a new user API implementation.
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
