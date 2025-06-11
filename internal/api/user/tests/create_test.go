package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	userApi "github.com/igorezka/auth/internal/api/user"
	"github.com/igorezka/auth/internal/model"
	"github.com/igorezka/auth/internal/service"
	serviceMocks "github.com/igorezka/auth/internal/service/mocks"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

func TestCreate(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		role     = gofakeit.IntRange(0, 1)
		password = gofakeit.Password(true, true, true, false, false, 16)

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			UserCreate: &desc.UserCreate{
				Info: &desc.UserInfo{
					Name:  name,
					Email: email,
					Role:  desc.Role(role),
				},
				Password:        password,
				PasswordConfirm: password,
			},
		}

		userCreate = &model.UserCreate{
			Name:     name,
			Email:    email,
			Role:     model.Role(role),
			Password: password,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, userCreate).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, userCreate).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)
			api := userApi.NewImplementation(userServiceMock)

			resHandler, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
