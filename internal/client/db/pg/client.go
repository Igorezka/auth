package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	def "github.com/igorezka/auth/internal/client/db"
)

type pgClient struct {
	masterDBC def.DB
}

// New constructor pgClient
func New(ctx context.Context, dsn string) (def.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, errors.Errorf("failed to connect to db: %v", err)
	}

	return &pgClient{
		masterDBC: NewDB(dbc),
	}, nil
}

func (c *pgClient) DB() def.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
