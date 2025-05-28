package accountservice

import (
	"encoding/json"
	"net/http"
)

func (service service) Delete(response http.ResponseWriter, request *http.Request) {
	accountID, err := GetIDFromRequest(request)
	response.Header().Add("Content-Type", "application/json")
	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]string{"error": err.Error()})
		
		return
	}

	err = service.usecase.Delete(accountID)
	if err != nil {
		response.WriteHeader(500)
		json.NewEncoder(response).Encode(map[string]string{"error": err.Error()})
		
		return
	}

	json.NewEncoder(response).Encode(map[string]any{
		"message": "Usuário deletado com sucesso",
		"userID":  accountID,
	},
	)
}
