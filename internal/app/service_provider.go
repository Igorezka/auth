package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/igorezka/auth/internal/api/user"
	"github.com/igorezka/auth/internal/closer"
	"github.com/igorezka/auth/internal/config"
	"github.com/igorezka/auth/internal/config/env"
	"github.com/igorezka/auth/internal/repository"
	userRepository "github.com/igorezka/auth/internal/repository/user"
	"github.com/igorezka/auth/internal/service"
	userService "github.com/igorezka/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       *pgxpool.Pool
	userRepository repository.UserRepository

	userService service.UserService
	userImpl    *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) *pgxpool.Pool {
	if s.dbClient == nil {
		cl, err := pgxpool.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %s", err.Error())
		}

		err = cl.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %s", err.Error())
		}
		closer.Add(func() error {
			cl.Close()
			return nil
		})

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
