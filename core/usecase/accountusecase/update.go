package accountusecase

import (
	"fmt"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/security"
)

func (usecase usecase) Update(updateAccountRequestDto *dto.UpdateAccountRequestDTO) (*domain.PublicAccount, error) {
	if updateAccountRequestDto.Password != nil {
		hashedPassword, err := security.HashPassword(*updateAccountRequestDto.Password)
		if err != nil {
			return nil, err
		}
		updateAccountRequestDto.Password = &hashedPassword
	}

	account, err := usecase.repository.GetAccountByID(updateAccountRequestDto.ID)
	if err != nil {
		return nil, fmt.Errorf("conta não encontrada")
	}

	if updateAccountRequestDto.FirstName != nil {
		account.FirstName = *updateAccountRequestDto.FirstName
	}
	if updateAccountRequestDto.LastName != nil {
		account.LastName = *updateAccountRequestDto.LastName
	}
	if updateAccountRequestDto.Password != nil {
		account.Password = *updateAccountRequestDto.Password
	}

	publicAccount, err := usecase.repository.Update(updateAccountRequestDto)

	if err != nil {
		return nil, err
	}

	return publicAccount, nil
}
