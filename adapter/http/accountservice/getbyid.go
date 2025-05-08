package accountservice

import (
	"encoding/json"
	"net/http"
)

func (service service) GetByID(response http.ResponseWriter, request *http.Request) {
	accountID, err := getID(request)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	account, err := service.usecase.GetByID(accountID)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(account)
}
