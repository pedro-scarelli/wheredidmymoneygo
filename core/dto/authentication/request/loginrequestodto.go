package dto

import (
	"encoding/json"
	"io"
)

type LoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FromJSONCreateLoginRequestDTO(body io.Reader) (*LoginRequestDTO, error) {
	CreateLoginRequestDTO := LoginRequestDTO{}
	if err := json.NewDecoder(body).Decode(&CreateLoginRequestDTO); err != nil {
		return nil, err
	}

	return &CreateLoginRequestDTO, nil
}
