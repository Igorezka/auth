package user

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/igorezka/auth/pkg/user_v1"
)

// Delete delete user
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("delete user with id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
