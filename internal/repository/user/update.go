package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/igorezka/auth/internal/model"
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

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
