package authenticationservice

import (
	"encoding/json"
	"net/http"

	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/authentication/request"
)

func (service service) Login(response http.ResponseWriter, request *http.Request) {
	accountRequest, err := dto.FromJSONCreateLoginRequestDTO(request.Body)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	account, err := service.usecase.Login(accountRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(response).Encode(account)
}
