package accountusecase

import (
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

	account, err := usecase.repository.Update(updateAccountRequestDto)

	if err != nil {
		return nil, err
	}

	return account, nil
}
