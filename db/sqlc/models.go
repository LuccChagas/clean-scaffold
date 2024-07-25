// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID      `json:"id"`
	Age       sql.NullInt32  `json:"age"`
	Name      sql.NullString `json:"name"`
	BirthDate sql.NullTime   `json:"birth_date"`
	Status    sql.NullBool   `json:"status"`
}
