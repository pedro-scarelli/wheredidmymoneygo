package accountservice

import (
	"encoding/json"
	"errors"
	"net/http"

	domain "github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (service service) GetByID(response http.ResponseWriter, request *http.Request) {
	accountID, err := GetIDFromRequest(request)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]string{"error": err.Error()})
		return
	}

	account, err := service.GetAccountByID(accountID)
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(map[string]string{"error": "account not found"})

			return
		}
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(map[string]string{"error": err.Error()})

		return
	}

	json.NewEncoder(response).Encode(account)
}

func (service *service) GetAccountByID(id string) (*domain.PublicAccount, error) {
	account, err := service.usecase.GetByID(id)

	return account, err
}
