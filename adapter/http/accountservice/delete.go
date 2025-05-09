package accountservice

import (
	"encoding/json"
	"net/http"
)

func (service service) Delete(response http.ResponseWriter, request *http.Request) {
	accountID, err := getID(request)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	err = service.usecase.Delete(accountID)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(map[string]any{
		"message": "Usuário deletado com sucesso",
		"userID":  accountID,
	},
	)
}
