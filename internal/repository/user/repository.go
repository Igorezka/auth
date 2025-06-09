package user

import (
	"github.com/igorezka/auth/internal/client/db"
	def "github.com/igorezka/auth/internal/repository"
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
