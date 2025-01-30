package middleware

import (
	"context"
	"ejTestTask/config"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

func AuthMiddleware(requiredRole string, next http.Handler) http.Handler {
	jwtKey := []byte(os.Getenv("jwtKey"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Missing token"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, `{"error": "Invalid token format"}`, http.StatusUnauthorized)
			return
		}

		claims := &config.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
			return
		}

		if requiredRole != "" && claims.Role != requiredRole {
			http.Error(w, `{"error": "Access denied"}`, http.StatusForbidden)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "username", claims.Username)
		ctx = context.WithValue(ctx, "role", claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GenerateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &config.Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("jwtKey")))
}
