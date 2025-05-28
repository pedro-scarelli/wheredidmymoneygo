package accountusecase

import (
	"time"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain/enum"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

func (usecase usecase) Movement(movementRequestDto *dto.MovementRequestDTO) error {
	if movementRequestDto.Type == enum.DEBITO {
		movementRequestDto.Value = -1 * movementRequestDto.Value
	}

	err := usecase.repository.Movement(movementRequestDto, time.Now().UTC())
	if err != nil {
		return err
	}

	return nil
}
