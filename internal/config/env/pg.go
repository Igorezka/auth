package env

import (
	"errors"
	"fmt"
	"os"

	def "github.com/igorezka/auth/internal/config"
)

var _ def.PGConfig = (*pgConfig)(nil)

const (
	pgHostEnvName     = "PG_HOST"
	pgPortEnvName     = "PG_PORT"
	pgDbnameEnvName   = "PG_DB"
	pgUserEnvName     = "PG_USER"
	pgPasswordEnvName = "PG_PASSWORD"
)

type pgConfig struct {
	dsn string
}

// NewPgConfig pg config struct constructor
func NewPgConfig() (*pgConfig, error) {
	host := os.Getenv(pgHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("pg host not found")
	}

	port := os.Getenv(pgPortEnvName)
	if len(host) == 0 {
		return nil, errors.New("pg port not found")
	}

	dbname := os.Getenv(pgDbnameEnvName)
	if len(host) == 0 {
		return nil, errors.New("pg dbname not found")
	}

	user := os.Getenv(pgUserEnvName)
	if len(host) == 0 {
		return nil, errors.New("pg user not found")
	}

	password := os.Getenv(pgPasswordEnvName)
	if len(host) == 0 {
		return nil, errors.New("pg password not found")
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
