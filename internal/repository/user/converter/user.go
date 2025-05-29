package converter

import (
	"github.com/igorezka/auth/internal/model"
	modelRepo "github.com/igorezka/auth/internal/repository/user/model"
)

// ToUserFromRepo converts a User repository model to business logic model.
func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToUserInfoFromRepo converts a UserInfo repository model to business logic model.
func ToUserInfoFromRepo(info modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  model.Role(info.Role),
	}
}

// ToUserCreateFromService converts a UserCreate business logic model to repository model.
func ToUserCreateFromService(userCreate *model.UserCreate) *modelRepo.UserCreate {
	return &modelRepo.UserCreate{
		Name:         userCreate.Name,
		Email:        userCreate.Email,
		Role:         modelRepo.Role(userCreate.Role),
		PasswordHash: []byte(userCreate.Password),
	}
}
