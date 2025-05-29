package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	domain "github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
    AccountID string `json:"sub"`
    jwt.RegisteredClaims
}

type contextKey string

const userClaimsKey contextKey = "userClaims"

func JwtAuthorizer(accountUseCase domain.AccountUseCase) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                respondUnauthorized(w)
                return
            }

            tokenParts := strings.Split(authHeader, " ")
            if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
                respondUnauthorized(w)
                return
            }

            tokenStr := tokenParts[1]
            claims := &Claims{}

            token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("wrong signature method: %v", token.Header["alg"])
                }
                return jwtSecret, nil
            })
            
            if err != nil || !token.Valid {
                respondUnauthorized(w)
                return
            }
            
            if claims.AccountID != "" {
                _, err := accountUseCase.GetByID(claims.AccountID)
                if err != nil {
                    if errors.Is(err, domain.ErrAccountNotFound) {
                        respondUnauthorized(w)
                        return
                    }
                    fmt.Printf("error checking account existence: %v\n", err)
                    respondWithError(w, http.StatusInternalServerError, "internal server error")
                    return
                }
            }
            
            ctx := context.WithValue(r.Context(), userClaimsKey, claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(fmt.Appendf(nil, `{"error": "%s"}`, message))
}

func respondUnauthorized(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(401)
    w.Write(fmt.Appendf(nil, `{"message": "%s"}`, "unauthorized"))
}