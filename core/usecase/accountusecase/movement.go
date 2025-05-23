package accountusecase

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
	"time"
)

func (usecase usecase) Movement(movementRequestDto *dto.MovementRequestDTO) (*domain.Movement, error) {
	movementCreated, err := usecase.repository.Movement(movementRequestDto, time.Now().UTC())
	if err != nil {
		return nil, err
	}

	return movementCreated, nil
}
