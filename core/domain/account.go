package domain

import (
	"net/http"

	"time"

	enum "github.com/pedro-scarelli/wheredidmymoneygo/core/domain/enum"
	accountRequestDto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

type PublicAccount struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Number       int64     `json:"number"`
	CPF          string    `json:"cpf"`
	Email        string    `json:"email"`
	MonthBalance float64   `json:"monthBalance"`
	CreatedAt    time.Time `json:"createdAt"`
}

type Account struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Number       int64     `json:"number"`
	CPF          string    `json:"cpf"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	MonthBalance int64     `json:"monthBalance"`
	CreatedAt    time.Time `json:"createdAt"`
	DeletedAt    time.Time `json:"deletedAt"`
}

type Movement struct {
	ID          string            `json:"id"`
	Type        enum.MovementType `json:"type"`
	Value       int               `json:"value"`
	DueDate     time.Time         `json:"dueDate"`
	AccountID   string            `json:"accountId"`
	Description string            `json:"description"`
	CreatedAt   time.Time         `json:"createdAt"`
}

type AccountService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
	Movement(response http.ResponseWriter, request *http.Request)
}

type AccountUseCase interface {
	Create(createAccountRequestDto *accountRequestDto.CreateAccountRequestDTO) (*PublicAccount, error)
	Delete(accountID string) error
	Update(updateAccountRequestDto *accountRequestDto.UpdateAccountRequestDTO) error
	GetByID(accountID string) (*PublicAccount, error)
	Movement(movementRequestDto *accountRequestDto.MovementRequestDTO) error
}

type AccountRepository interface {
	Create(createAccountRequestDto *accountRequestDto.CreateAccountRequestDTO, accountNumber int, createdAt time.Time) (*PublicAccount, error)
	Delete(accountID string) error
	Update(updateAccountRequestDto *accountRequestDto.UpdateAccountRequestDTO) error
	GetAccountByID(accountID string) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
	Movement(movementRequestDto *accountRequestDto.MovementRequestDTO, createdAt time.Time) error
	GetAccountBalance(accountID string) (int, error)
}
