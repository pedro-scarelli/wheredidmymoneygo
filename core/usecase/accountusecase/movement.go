package accountusecase

import (
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

func (usecase usecase) Movement(movementRequestDto *dto.MovementRequestDTO) error {
	err := usecase.repository.Movement(movementRequestDto)
	if err != nil {
		return err
	}

	return nil
}
