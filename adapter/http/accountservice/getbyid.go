package accountservice

import (
	"encoding/json"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"net/http"
)

func (service service) GetByID(response http.ResponseWriter, request *http.Request) {
	accountID, err := GetIDFromRequest(request)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))

		return
	}

	account, err := service.GetAccountByID(accountID)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))

		return
	}

	json.NewEncoder(response).Encode(account)
}

func (service *service) GetAccountByID(id string) (*domain.PublicAccount, error) {
	account, err := service.usecase.GetByID(id)

	return account, err
}
