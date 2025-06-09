package tests

import (
	"context"
	"fmt"
	"github.com/igorezka/auth/internal/model"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"

	"github.com/igorezka/auth/internal/service"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

func TestCreate(t *testing.T) {
	type userServiceMockFunc func(ms *minimock.Controller) service.UserService

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
		role     = desc.Role_user
		password = gofakeit.Password(true, true, true, false, false, 3)

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			UserCreate: &desc.UserCreate{
				Info: &desc.UserInfo{
					Name:  name,
					Email: email,
					Role:  role,
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
		{},
	}
}
