package dto

type UpdateAccountRequest struct {
	ID        int64   `json:"id"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Password  *string `json:"password,omitempty"`
}
