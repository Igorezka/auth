package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/igorezka/auth/internal/client/db"
	dbClientMocks "github.com/igorezka/auth/internal/client/db/mocks"
	"github.com/igorezka/auth/internal/model"
	"github.com/igorezka/auth/internal/repository"
	repositoryMocks "github.com/igorezka/auth/internal/repository/mocks"
	userService "github.com/igorezka/auth/internal/service/user"
)

func TestGet(t *testing.T) {
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMockFunc func(f func(ctx context.Context) error, mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		role      = gofakeit.IntRange(0, 1)
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		repoErr = fmt.Errorf("repository error")

		user = &model.User{
			ID: id,
			Info: model.UserInfo{
				Name:  name,
				Email: email,
				Role:  model.Role(role),
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{Time: updatedAt, Valid: true},
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMockFunc
		txManagerMock      txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: user,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(user, nil)
				return mock
			},
			txManagerMock: func(_ func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := dbClientMocks.NewTxManagerMock(mc)
				return mock
			},
		},
		{
			name: "user repository error",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repoErr)
				return mock
			},
			txManagerMock: func(_ func(context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := dbClientMocks.NewTxManagerMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			userRepositoryMock := tt.userRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(func(_ context.Context) error {
				return nil
			}, mc)
			service := userService.NewService(userRepositoryMock, txManagerMock)

			resHandler, err := service.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
