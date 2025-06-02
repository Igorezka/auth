package user

import (
	"github.com/igorezka/auth/internal/repository"
	def "github.com/igorezka/auth/internal/service"
)

var _ def.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
}

// NewService creates a new user service.
func NewService(userRepository repository.UserRepository) *serv {
	return &serv{
		userRepository: userRepository,
	}
}
