package accountservice

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/http/middleware"
)

func GetIDFromRequest(r *http.Request) (string, error) {
	id := mux.Vars(r)["id"]

	return id, nil
}

func GetAccountIDFromToken(r *http.Request) (string, error) {
	claims, ok := r.Context().Value(middleware.UserClaimsKey).(*middleware.Claims)
	if !ok {
		return "", fmt.Errorf("erro ao obter claims do token")
	}

	return claims.AccountID, nil
}
