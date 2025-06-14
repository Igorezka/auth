package user

import (
	"context"

	"github.com/igorezka/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, userCreate *model.UserCreate) (int64, error) {
	id, err := s.userRepository.Create(ctx, userCreate)
	if err != nil {
		return 0, err
	}

	return id, nil
}
