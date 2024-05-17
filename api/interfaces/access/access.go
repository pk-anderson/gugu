package access

import "time"

type Access struct {
	AccessID    string
	UserID      string
	AccessToken string
	ExpiresAt   time.Time
	Revoked     bool
	CreatedAt   time.Time
	SessionID   string
}
