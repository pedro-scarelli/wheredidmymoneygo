package accountusecase

import (
	"errors"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (usecase usecase) GetByID(accountID string) (*domain.PublicAccount, error) {
	account, err := usecase.repository.GetAccountByID(accountID)
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			return nil, err
		}
		return nil, errors.New("failed to get account")
	}
	
	balance, err := usecase.repository.GetAccountBalance(accountID)
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			return nil, err
		}
		return nil, errors.New("failed to get account movements")
	}


	publicAccount := &domain.PublicAccount{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Number:    account.Number,
		CPF:       account.CPF,
		Email:     account.Email,
		Balance:   float64(balance) / 100,
		CreatedAt: account.CreatedAt,
	}

	return publicAccount, nil
}
