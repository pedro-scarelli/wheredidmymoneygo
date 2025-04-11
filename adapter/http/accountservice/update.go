package userservice

import (
	"encoding/json"
	"net/http"

	"github.com/pedro-scarelli/go_login/core/dto"
)

func (service service) Update(response http.ResponseWriter, request *http.Request) {
	userRequest, err := dto.FromJSONCreateUserRequest(request.Body)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	user, err := service.usecase.Create(userRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(user)
}
