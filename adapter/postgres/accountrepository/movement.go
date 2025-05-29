package accountrepository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

func (repository repository) Movement(movementRequestDto *dto.MovementRequestDTO, createdAt time.Time) error {
	ctx := context.Background()

	recurrence := movementRequestDto.Recurrence
	ids := make([]string, recurrence)
	dueDates := make([]time.Time, recurrence)
	values := make([]int, recurrence)

	for i := range recurrence {
		ids[i] = uuid.New().String()
		dueDates[i] = movementRequestDto.DueDate.AddDate(0, i, 0)
		values[i] = int(movementRequestDto.Value * 100)
	}

	query := `
        INSERT INTO tb_movement 
            (pk_st_id, it_value, dt_due_date, st_account_id, st_description, dt_created_at)
        SELECT 
            unnest($1::uuid[]), 
            unnest($2::integer[]), 
            unnest($3::date[]),
            unnest($4::text[]), 
            unnest($5::text[]), 
            unnest($6::timestamp[])
    `

	_, err := repository.db.Exec(
		ctx,
		query,
		ids,
		values,
		dueDates,
		RepeatString(movementRequestDto.AccountID, recurrence),
		RepeatString(movementRequestDto.Description, recurrence),
		RepeatTime(createdAt, recurrence),
	)

	if err != nil {
		return fmt.Errorf("erro ao inserir movimentos: %w", err)
	}

	return nil
}

func RepeatString(value string, n int) []string {
	arr := make([]string, n)
	for i := range arr {
		arr[i] = value
	}
	return arr
}

func RepeatTime(value time.Time, n int) []time.Time {
	arr := make([]time.Time, n)
	for i := range arr {
		arr[i] = value
	}
	return arr
}
