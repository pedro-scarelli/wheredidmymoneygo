package dto

import (
	"encoding/json"
	"io"
)

type UpdateAccountRequestDTO struct {
	ID        string  `json:"id"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Password  *string `json:"password,omitempty"`
}

func FromJSONUpdateAccountRequestDTO(body io.Reader) (*UpdateAccountRequestDTO, error) {
	updateAccountRequestDto := UpdateAccountRequestDTO{}
	if err := json.NewDecoder(body).Decode(&updateAccountRequestDto); err != nil {
		return nil, err
	}

	return &updateAccountRequestDto, nil
}
