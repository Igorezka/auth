package user

import (
	"context"

	"github.com/igorezka/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, userUpdate *model.UserUpdate) error {
	err := s.userRepository.Update(ctx, id, userUpdate)
	if err != nil {
		return err
	}

	return nil
}
