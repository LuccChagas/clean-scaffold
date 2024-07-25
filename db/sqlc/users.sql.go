// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, age, name, birth_date, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, age, name, birth_date, status
`

type CreateUserParams struct {
	ID        uuid.UUID      `json:"id"`
	Age       sql.NullInt32  `json:"age"`
	Name      sql.NullString `json:"name"`
	BirthDate sql.NullTime   `json:"birth_date"`
	Status    sql.NullBool   `json:"status"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Age,
		arg.Name,
		arg.BirthDate,
		arg.Status,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Age,
		&i.Name,
		&i.BirthDate,
		&i.Status,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, age, name, birth_date, status FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Age,
			&i.Name,
			&i.BirthDate,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersByID = `-- name: GetUsersByID :one
SELECT id, age, name, birth_date, status FROM users WHERE id = $1
`

func (q *Queries) GetUsersByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsersByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Age,
		&i.Name,
		&i.BirthDate,
		&i.Status,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET age = $2, name = $3, birth_date = $4, status = $5
WHERE id = $1
RETURNING id, age, name, birth_date, status
`

type UpdateUserParams struct {
	ID        uuid.UUID      `json:"id"`
	Age       sql.NullInt32  `json:"age"`
	Name      sql.NullString `json:"name"`
	BirthDate sql.NullTime   `json:"birth_date"`
	Status    sql.NullBool   `json:"status"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Age,
		arg.Name,
		arg.BirthDate,
		arg.Status,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Age,
		&i.Name,
		&i.BirthDate,
		&i.Status,
	)
	return i, err
}