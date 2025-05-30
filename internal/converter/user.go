package converter

import (
	"github.com/igorezka/auth/internal/model"
	desc "github.com/igorezka/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
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

// ToUserUpdateFromDesc converts the UserUpdate model from proto-model to a business logic model.
func ToUserUpdateFromDesc(userUpdate *desc.UserUpdate) *model.UserUpdate {
	var (
		name  *string
		email *string
		role  model.Role
	)

	if userUpdate.GetName() != nil {
		name = &userUpdate.GetName().Value
	}
	if userUpdate.GetEmail() != nil {
		email = &userUpdate.GetEmail().Value
	}
	role = (model.Role)(userUpdate.GetRole())

	return &model.UserUpdate{
		Name:  name,
		Email: email,
		Role:  role,
	}
}

// ToUserFromService converts a User service model to proto-model.
func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id: user.ID,
		Info: &desc.UserInfo{
			Name:  user.Info.Name,
			Email: user.Info.Email,
			Role:  desc.Role(user.Info.Role),
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
