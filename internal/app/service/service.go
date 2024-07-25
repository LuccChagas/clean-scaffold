package service

import (
	"context"
	"github.com/LuccChagas/clean-scaffold/internal/app/model"
)

type UserServiceInterface interface {
	CreateUserService(ctx context.Context, user model.User) (response model.User, err error)
}
