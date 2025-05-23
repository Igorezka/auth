package config

import (
	"github.com/joho/godotenv"
)

// GRPCConfig is interface containing gpc config methods
type GRPCConfig interface {
	Address() string
}

// PGConfig is interface containing pg config methods
type PGConfig interface {
	DSN() string
}

// Load function for loading a configuration file from a given path
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
