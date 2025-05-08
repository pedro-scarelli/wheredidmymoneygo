package accountusecase

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (usecase usecase) GetByID(accountID int) (*domain.Account, error) {
	account, err := usecase.repository.GetByID(accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
