package user

import (
	"context"
	"log"

	"github.com/igorezka/auth/internal/converter"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserCreateFromDesc(req.GetUserCreate()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
