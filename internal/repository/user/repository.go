package user

import (
	def "github.com/igorezka/auth/internal/repository"
	"github.com/igorezka/zdb_platform_common/pkg/client/db"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates new user repository
func NewRepository(db db.Client) def.UserRepository {
	return &repo{db: db}
}
