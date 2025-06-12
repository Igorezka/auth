package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/igorezka/auth/internal/model"
	"github.com/igorezka/auth/internal/repository"
	repositoryMocks "github.com/igorezka/auth/internal/repository/mocks"
	userService "github.com/igorezka/auth/internal/service/user"
	"github.com/igorezka/zdb_platform_common/pkg/client/db"
	dbClientMocks "github.com/igorezka/zdb_platform_common/pkg/client/db/mocks"
)

func TestUpdate(t *testing.T) {
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type txManagerMockFunc func(f func(ctx context.Context) error, mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		id  int64
		req *model.UserUpdate
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = gofakeit.IntRange(0, 1)

		repoErr = fmt.Errorf("repository error")

		req = &model.UserUpdate{
			Name:  &name,
			Email: &email,
			Role:  model.Role(role),
		}
	)

	tests := []struct {
		name               string
		args               args
		err                error
		userRepositoryMock userRepositoryMockFunc
		txManagerMock      txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				id:  id,
				req: req,
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, id, req).Return(nil)
				return mock
			},
			txManagerMock: func(_ func(ctx context.Context) error, mc *minimock.Controller) db.TxManager {
				mock := dbClientMocks.NewTxManagerMock(mc)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				id:  id,
				req: req,
			},
			err: repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, id, req).Return(repoErr)
				return mock
			},
			txManagerMock: func(_ func(ctx context.Context) error, mc *minimock.Controller) db.TxManager {
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

			err := service.Update(tt.args.ctx, tt.args.id, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
