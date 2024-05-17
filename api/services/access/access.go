package access

import (
	"database/sql"
	"errors"
	access "gugu/interfaces/access"
	"gugu/repositories/accessRepository"
	"gugu/repositories/userRepository"
	"gugu/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type AccessService interface {
	Login(email, password string) (string, error)
}

type accessService struct {
	DB *sql.DB
}

type Claims struct {
	UserID   string `json:"userId"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *accessService) Login(email, password string) (string, error) {
	userRep := userRepository.NewRepository(s.DB)
	user, err := userRep.VerifyCredentials(email, password)
	if err != nil {
		return "", err
	}

	err = godotenv.Load()
	if err != nil {
		return "", errors.New("error on reading jwt secret")
	}
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.UserId,
		Email:    user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	accessRep := accessRepository.NewRepository(s.DB)
	access := &access.Access{
		AccessID:    utils.GenerateUUID(),
		UserID:      user.UserId,
		ExpiresAt:   expirationTime,
		AccessToken: token,
		Revoked:     false,
		CreatedAt:   time.Now(),
		SessionID:   utils.GenerateUUID(),
	}

	err = accessRep.InsertAccessToken(access)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewService(DB *sql.DB) AccessService {
	return &accessService{
		DB: DB,
	}
}
