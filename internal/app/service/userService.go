package service

import (
	"context"
	"database/sql"
	db "github.com/LuccChagas/clean-scaffold/db/sqlc"
	"github.com/LuccChagas/clean-scaffold/internal/app/model"
	"github.com/LuccChagas/clean-scaffold/internal/app/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repo repository.RepositoryInterface
}

func NewUserService(repo repository.RepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUserService(ctx context.Context, user model.User) (response model.User, err error) {
	arg := db.CreateUserParams{
		ID: uuid.UUID{},
		Age: sql.NullInt32{
			Int32: int32(user.Age),
			Valid: true,
		},
		Name: sql.NullString{
			String: user.Name,
			Valid:  true,
		},
		BirthDate: sql.NullTime{
			Time:  user.BirthDate,
			Valid: true,
		},
		Status: sql.NullBool{
			Bool:  user.Status,
			Valid: true,
		},
	}

	u, err := s.repo.CreateUserRepo(ctx, arg)
	if err != nil {
		return response, err
	}

	response = model.User{
		ID:        u.ID,
		Age:       uint32(u.Age.Int32),
		Name:      u.Name.String,
		BirthDate: u.BirthDate.Time,
		Status:    u.Status.Bool,
	}

	return response, err
}
