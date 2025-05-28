package user

import (
	"context"
	"github.com/igorezka/auth/internal/repository"
	desc "github.com/igorezka/auth/pkg/user_v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, user *desc.CreateRequest) (int64, error) {
	return 0, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*desc.User, error) {
	return nil, nil
}

func (r *repo) Update(ctx context.Context, user *desc.UpdateRequest) error {
	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	return nil
}
