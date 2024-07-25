package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Age       uint32    `json:"age,omitempty"`
	Name      string    `json:"name,omitempty"`
	BirthDate time.Time `json:"birth_date"`
	Status    bool      `json:"status"`
}
