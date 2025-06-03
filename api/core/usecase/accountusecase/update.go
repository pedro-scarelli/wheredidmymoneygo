package accountusecase

import (
	"fmt"

	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/security"
)

func (usecase usecase) Update(updateAccountRequestDto *dto.UpdateAccountRequestDTO) error {
	if updateAccountRequestDto.Password != nil {
		hashedPassword, err := security.HashPassword(*updateAccountRequestDto.Password)
		if err != nil {
			return err
		}
		updateAccountRequestDto.Password = &hashedPassword
	}

	account, err := usecase.repository.GetAccountByID(updateAccountRequestDto.ID)
	if err != nil {
		return fmt.Errorf("conta não encontrada")
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

	err = usecase.repository.Update(updateAccountRequestDto)

	if err != nil {
		return err
	}

	return nil
}
