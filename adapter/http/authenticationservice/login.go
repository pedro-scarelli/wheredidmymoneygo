package authenticationservice

import (
	"encoding/json"
	"net/http"

	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/authentication/request"
)

func (service service) Login(response http.ResponseWriter, request *http.Request) {
	loginRequestDto, err := dto.FromJSONCreateLoginRequestDTO(request.Body)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	loginResponseDto, err := service.usecase.Login(loginRequestDto)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(response).Encode(loginResponseDto)
}
