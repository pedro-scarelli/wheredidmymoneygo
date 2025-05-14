package domain

import (
	"net/http"

	"time"

	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
	accountRequestDto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

type PublicAccount struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type Account struct {
	ID        string    `json:"id"`
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
	Create(accountRequest *accountRequestDto.CreateAccountRequestDTO) (*PublicAccount, error)
	Delete(accountID string) error
	Update(accountRequest *accountRequestDto.UpdateAccountRequestDTO) (*PublicAccount, error)
	Get(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]PublicAccount], error)
	GetByID(accountID string) (*PublicAccount, error)
}

type AccountRepository interface {
	Create(accountRequest *accountRequestDto.CreateAccountRequestDTO, accountNumber int, createdAt time.Time) (*PublicAccount, error)
	Delete(accountID string) error
	Update(accountRequest *accountRequestDto.UpdateAccountRequestDTO) (*PublicAccount, error)
	Get(paginationRequestParams *dto.PaginationRequestParams) (*Pagination[[]PublicAccount], error)
	GetAccountByID(accountID string) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
}
