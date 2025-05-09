package domain

import (
	"net/http"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
	"time"
)

type PublicAccount struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type AccountService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
}

type AccountUseCase interface {
	Create(accountRequest *dto.CreateAccountRequest) (*PublicAccount, error)
	Delete(accountID int) error
	Update(accountRequest *dto.UpdateAccountRequest) (*PublicAccount, error)
	Get(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]PublicAccount], error)
	GetByID(accountID int) (*PublicAccount, error)
}

type AccountRepository interface {
	Create(accountRequest *dto.CreateAccountRequest, accountNumber int, createdAt time.Time) (*PublicAccount, error)
	Delete(accountID int) error
	Update(accountRequest *dto.UpdateAccountRequest) (*PublicAccount, error)
	Get(paginationRequestParams *dto.PaginationRequestParams) (*Pagination[[]PublicAccount], error)
	GetByID(accountID int) (*PublicAccount, error)
}
