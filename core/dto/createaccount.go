package dto

import (
	"encoding/json"
	"io"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func FromJSONCreateAccountRequest(body io.Reader) (*CreateAccountRequest, error) {
	createAccountRequest := CreateAccountRequest{}
	if err := json.NewDecoder(body).Decode(&createAccountRequest); err != nil {
		return nil, err
	}

	return &createAccountRequest, nil
}
