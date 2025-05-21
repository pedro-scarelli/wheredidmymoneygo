package accountservice

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/http/middleware"
	"net/http"
)

func GetIDFromRequest(r *http.Request) (string, error) {
	id := mux.Vars(r)["id"]

	return id, nil
}

func GetUserIDFromToken(r *http.Request) (string, error) {
	claims, ok := r.Context().Value("userClaims").(*middleware.Claims)
	if !ok {
		return "", fmt.Errorf("Erro ao obter claims do token")
	}

	return claims.UserID, nil
}
