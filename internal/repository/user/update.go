package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/igorezka/auth/internal/model"
	"github.com/igorezka/zdb_platform_common/pkg/client/db"
)

func (r *repo) Update(ctx context.Context, id int64, userUpdate *model.UserUpdate) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar)

	if userUpdate.Name != nil {
		builder = builder.Set(nameColumn, userUpdate.Name)
	}

	if userUpdate.Email != nil {
		builder = builder.Set(emailColumn, userUpdate.Email)
	}

	builder = builder.Set(roleColumn, userUpdate.Role)

	builder = builder.Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
