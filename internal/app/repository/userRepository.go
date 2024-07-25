package repository

import (
	"context"
	db "github.com/LuccChagas/clean-scaffold/db/sqlc"
	"github.com/google/uuid"
)

func (r *Repository) CreateUserRepo(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.Queries.CreateUser(ctx, arg)
}

func (r *Repository) GetAllUsersRepo(ctx context.Context) ([]db.User, error) {
	return r.Queries.GetAllUsers(ctx)
}

func (r *Repository) UpdateUserRepo(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.Queries.UpdateUser(ctx, arg)
}

func (r *Repository) GetUsersByIDRepo(ctx context.Context, u uuid.UUID) (db.User, error) {
	return r.Queries.GetUsersByID(ctx, u)
}
