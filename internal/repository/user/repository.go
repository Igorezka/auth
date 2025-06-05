package user

import (
	"github.com/igorezka/auth/internal/client/db"
	def "github.com/igorezka/auth/internal/repository"
)

var _ def.UserRepository = (*repo)(nil)

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
func NewRepository(db db.Client) *repo {
	return &repo{db: db}
}
