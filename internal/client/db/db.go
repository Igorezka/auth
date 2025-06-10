package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Handler function in tx manager
type Handler func(ctx context.Context) error

// Client is a client to work with the database
type Client interface {
	DB() DB
	Close() error
}

// TxManager interface to manage transactions
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query wrapper over the request, contains the name of the query and the query itself
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor interface for work with transactions
type Transactor interface {
	BeginTx(ctx context.Context, tx pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecer combines NamedExecer and QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer interface to work with the named query using tags in structures
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer interface to work with the basic query
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger interface to check db connection
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB interface for working with the database
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
