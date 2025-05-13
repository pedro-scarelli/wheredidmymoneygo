package accountservice

import (
	"encoding/json"
	"net/http"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
)

func (service service) Update(response http.ResponseWriter, request *http.Request) {
	accountRequest, err := dto.FromJSONUpdateAccountRequest(request.Body)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	account, err := service.usecase.Update(accountRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(account)
}
