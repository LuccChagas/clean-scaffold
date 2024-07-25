package repository

import (
	"context"
	db "github.com/LuccChagas/clean-scaffold/db/sqlc"
	"github.com/google/uuid"
)

type Repository struct {
	dbtx    db.DBTX
	Queries *db.Queries
}

func NewRepository(dbtx db.DBTX, q *db.Queries) *Repository {
	return &Repository{
		dbtx:    dbtx,
		Queries: q,
	}
}

// RepositoryInterface Add here all repository in one unique interface in order to share the calls easy in any service
type RepositoryInterface interface {
	CreateUserRepo(context.Context, db.CreateUserParams) (db.User, error)
	GetAllUsersRepo(context.Context) ([]db.User, error)
	GetUsersByIDRepo(context.Context, uuid.UUID) (db.User, error)
	UpdateUserRepo(context.Context, db.UpdateUserParams) (db.User, error)
}
