package domain

import (
	"net/http"

	"github.com/pedro-scarelli/go_login/core/dto"
	"time"
)

type PublicUser struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
}

type UserUseCase interface {
	Create(userRequest *dto.CreateUserRequest) (*PublicUser, error)
	Delete(userID int) error
	Update(userRequest *dto.CreateUserRequest) (*PublicUser, error)
	Get(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]User], error)
	GetByID(userID int) (*User, error)
}

type UserRepository interface {
	Create(userRequest *dto.CreateUserRequest, userNumber int, createdAt time.Time) (*PublicUser, error)
	Delete(userID int) error
	Update(userRequest *dto.CreateUserRequest) (*PublicUser, error)
	Get(paginationRequestParams *dto.PaginationRequestParms) (*Pagination[[]User], error)
	GetByID(userID int) (*User, error)
}
