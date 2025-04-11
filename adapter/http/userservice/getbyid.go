package userservice

import (
	"encoding/json"
	"net/http"
)

func (service service) GetByID(response http.ResponseWriter, request *http.Request) {
	userID, err := getID(request)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	user, err := service.usecase.GetByID(userID)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(user)
}
