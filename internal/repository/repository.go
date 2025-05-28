package repository

import (
	"context"

	"github.com/igorezka/auth/internal/repository/user/model"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

type UserRepository interface {
	Create(ctx context.Context, userCreate *model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*desc.User, error)
	Update(ctx context.Context, user *desc.UpdateRequest) error
	Delete(ctx context.Context, id int64) error
}
