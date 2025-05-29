package converter

import (
	"github.com/igorezka/auth/internal/model"
	desc "github.com/igorezka/auth/pkg/user_v1"
)

// ToUserCreateFromDesc converts the UserCreate model from proto-model to a business logic model.
func ToUserCreateFromDesc(userCreate *desc.UserCreate) *model.UserCreate {
	return &model.UserCreate{
		Name:     userCreate.GetInfo().GetName(),
		Email:    userCreate.GetInfo().GetEmail(),
		Role:     model.Role(userCreate.GetInfo().GetRole()),
		Password: userCreate.GetPassword(),
	}
}
