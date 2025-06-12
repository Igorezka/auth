package user

import (
	"github.com/igorezka/auth/internal/repository"
	def "github.com/igorezka/auth/internal/service"
	"github.com/igorezka/zdb_platform_common/pkg/client/db"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

// NewService creates a new user service.
func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) def.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
