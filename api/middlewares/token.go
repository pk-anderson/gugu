package middlewares

import (
	"context"
	"database/sql"
	access "gugu/interfaces/access"
	"gugu/repositories/accessRepository"
	"gugu/utils"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func AuthMiddleware(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := godotenv.Load(); err != nil {
			http.Error(w, "Error on reading jwt secret", http.StatusInternalServerError)
			return
		}
		jwtKey := []byte(os.Getenv("JWT_SECRET"))

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		claims := &access.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		rep := accessRepository.NewRepository(db)
		_, err = rep.GetAccessByToken(tokenString)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
			return
		}

		ctx := context.WithValue(r.Context(), utils.UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
