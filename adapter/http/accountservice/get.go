package accountservice

import (
	"encoding/json"
	"net/http"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
)

func (service service) Get(response http.ResponseWriter, request *http.Request) {
	paginationRequest, err := dto.FromValuePaginationRequestParams(request)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	accounts, err := service.usecase.Get(paginationRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(accounts)
}
