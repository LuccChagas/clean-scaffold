package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("token has expired")
var ErrInvalidToken = errors.New("token is invalid")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	UserID    string    `json:"user_id"`
	AccessKey int64     `json:"access_key"`
	AccessID  int64     `json:"access_id"`
	TenantID  string    `json:"tenant_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time
}

// NewPayload - TODO: Add what you need to generate a token payload - the under params are only example
func NewPayload(userID, username, tenantID string, accessKey int64, duration time.Duration, accessID int64) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		UserID:    userID,
		AccessKey: accessKey,
		AccessID:  accessID,
		TenantID:  tenantID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

func (payload *Payload) valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
