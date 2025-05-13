package dto

import (
	"encoding/json"
	"io"
)

type CreateAccountRequestDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func FromJSONCreateAccountRequestDTO(body io.Reader) (*CreateAccountRequestDTO, error) {
	CreateAccountRequestDTO := CreateAccountRequestDTO{}
	if err := json.NewDecoder(body).Decode(&CreateAccountRequestDTO); err != nil {
		return nil, err
	}

	return &CreateAccountRequestDTO, nil
}