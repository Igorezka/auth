package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/igorezka/auth/internal/client/db"
	"github.com/igorezka/auth/internal/model"
	"github.com/igorezka/auth/internal/repository/user/converter"
	modelRepo "github.com/igorezka/auth/internal/repository/user/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
