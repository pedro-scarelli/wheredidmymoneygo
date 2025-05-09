package accountusecase

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/security"
)

func (usecase usecase) Update(updateAccountRequest *dto.UpdateAccountRequest) (*domain.PublicAccount, error) {
	if updateAccountRequest.Password != nil {
		hashedPassword, err := security.HashPassword(*updateAccountRequest.Password)
		if err != nil {
			return nil, err
		}
		updateAccountRequest.Password = &hashedPassword
	}

	account, err := usecase.repository.Update(updateAccountRequest)

	if err != nil {
		return nil, err
	}

	return account, nil
}
