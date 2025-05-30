package user

import (
	"context"

	"github.com/igorezka/auth/internal/model"
)

func (s *serv) Get(ctx context.Context, userId int64) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
