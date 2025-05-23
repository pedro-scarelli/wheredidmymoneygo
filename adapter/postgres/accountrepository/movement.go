package accountrepository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

func (repository repository) Movement(movementRequestDto *dto.MovementRequestDTO, createdAt time.Time) (*domain.Movement, error) {
	ctx := context.Background()
	movement := domain.Movement{}

	err := repository.db.QueryRow(
		ctx,
		`INSERT INTO tb_movement
		(pk_st_id, it_value, it_recurrence, st_type, st_account_id, st_description, dt_created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING pk_st_id, it_value, it_recurrence, st_type, st_account_id, st_description, dt_created_at`,
		uuid.New().String(),
		movementRequestDto.Value*100,
		movementRequestDto.Recurrence,
		movementRequestDto.Type,
		movementRequestDto.AccountID,
		movementRequestDto.Description,
		createdAt,
	).Scan(
		&movement.ID,
		&movement.Value,
		&movement.Recurrence,
		&movement.Type,
		&movement.AccountID,
		&movement.Description,
		&movement.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &movement, nil
}
