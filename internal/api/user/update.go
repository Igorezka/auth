package user

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/igorezka/auth/internal/converter"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

// Update updates a user
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, req.GetId(), converter.ToUserUpdateFromDesc(req.GetUserUpdate()))
	if err != nil {
		return nil, err
	}

	log.Printf("updated user with id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
