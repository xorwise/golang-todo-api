package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/xorwise/golang-todo-api/domain"
)

type JWTMiddleware struct {
	Secret     string
	Repository domain.UserRepository
}

type userContextKey string

const (
	UserIDKey    userContextKey = "ID"
	UserEmailKey userContextKey = "Email"
)

func (j *JWTMiddleware) LoginRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &domain.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(j.Secret), nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*domain.JwtCustomClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		id := claims.ID
		_, err = j.Repository.GetByID(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		email := claims.Email

		ctx := context.WithValue(r.Context(), UserIDKey, id)
		ctx = context.WithValue(ctx, UserEmailKey, email)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	})
}
