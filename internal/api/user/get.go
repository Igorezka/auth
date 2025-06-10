package user

import (
	"context"
	"log"

	"github.com/igorezka/auth/internal/converter"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

// Get gets a user.
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("user email: %s", user.Info.Email)

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
