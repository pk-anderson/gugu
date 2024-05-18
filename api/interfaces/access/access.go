package access

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Access struct {
	AccessID    string
	UserID      string
	AccessToken string
	ExpiresAt   time.Time
	Revoked     bool
	CreatedAt   time.Time
	SessionID   string
}

type Claims struct {
	UserID   string `json:"userId"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
