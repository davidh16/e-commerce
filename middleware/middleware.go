package middleware

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

type Middleware struct {
	redis *redis.Client
}

func InitializeMiddleware(redis *redis.Client) *Middleware {
	return &Middleware{
		redis: redis,
	}
}

func (m *Middleware) AdminAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		token := strings.Split(authHeader, "Bearer ")[1]

		tkn, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
			return nil, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var role string
		if claims, ok := tkn.Claims.(jwt.MapClaims); ok {
			role = fmt.Sprint(claims["role"])
		}

		if role != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func (m *Middleware) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		token := strings.Split(authHeader, "Bearer ")[1]

		tkn, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
			return nil, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
