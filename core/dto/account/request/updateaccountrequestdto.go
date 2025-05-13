package dto

import (
	"encoding/json"
	"io"
)

type UpdateAccountRequestDTO struct {
	ID        int    `json:"id"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Password  *string `json:"password,omitempty"`
}

func FromJSONUpdateAccountRequestDTO(body io.Reader) (*UpdateAccountRequestDTO, error) {
	request := &UpdateAccountRequestDTO{}
	err := json.NewDecoder(body).Decode(request)

	return request, err
}
