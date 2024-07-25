package token

import (
	"time"
)

type Maker interface {
	CreateToken(userID, username, tenantID string, accessKey int64, duration time.Duration, accessID int64) (string, error)
	VerifyToken(token string) (*Payload, error)
}
