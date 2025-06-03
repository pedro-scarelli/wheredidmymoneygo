package accountservice

import (
	"encoding/json"
	"net/http"

	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	accountRequest, err := dto.FromJSONCreateAccountRequestDTO(request.Body)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]string{"error": err.Error()})

		return
	}

	account, err := service.usecase.Create(accountRequest)

	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]string{"error": err.Error()})
		
		return
	}
	json.NewEncoder(response).Encode(account)
}
