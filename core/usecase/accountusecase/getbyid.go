package accountusecase

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (usecase usecase) GetByID(accountID int) (*domain.PublicAccount, error) {
	account, err := usecase.repository.GetByID(accountID)
	if err != nil {
		return nil, err
	}

	publicAccount := &domain.PublicAccount{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Number:    account.Number,
		CPF:       account.CPF,
		Email:     account.Email,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}

	return publicAccount, nil
}
