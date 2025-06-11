package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	userApi "github.com/igorezka/auth/internal/api/user"
	"github.com/igorezka/auth/internal/model"
	"github.com/igorezka/auth/internal/service"
	serviceMocks "github.com/igorezka/auth/internal/service/mocks"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

func TestUpdate(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = gofakeit.IntRange(0, 1)

		serviceErr = fmt.Errorf("service error")

		req = &desc.UpdateRequest{
			Id: id,
			UserUpdate: &desc.UserUpdate{
				Name:  wrapperspb.String(name),
				Email: wrapperspb.String(email),
				Role:  desc.Role(role),
			},
		}

		userUpdate = &model.UserUpdate{
			Name:  &name,
			Email: &email,
			Role:  model.Role(role),
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &emptypb.Empty{},
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, req.Id, userUpdate).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, req.Id, userUpdate).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			api := userApi.NewImplementation(userServiceMock)

			resHandler, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
