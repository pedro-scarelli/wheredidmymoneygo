package dto

import (
	"encoding/json"
	"io"
)

type UpdateAccountRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

func FromJSONUpdateAccountRequest(body io.Reader) (*UpdateAccountRequest, error) {
	request := &UpdateAccountRequest{}
	err := json.NewDecoder(body).Decode(request)

	return request, err
}
