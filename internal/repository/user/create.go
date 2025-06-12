package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/igorezka/auth/internal/model"
	"github.com/igorezka/auth/internal/repository/user/converter"
	"github.com/igorezka/zdb_platform_common/pkg/client/db"
)

func (r *repo) Create(ctx context.Context, userCreate *model.UserCreate) (int64, error) {
	userCreateRepo := converter.ToUserCreateFromService(userCreate)

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn, passwordColumn).
		Values(userCreateRepo.Name, userCreateRepo.Email, userCreateRepo.Role, userCreateRepo.PasswordHash).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
