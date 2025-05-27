package accountrepository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
	"time"
)

func (repository repository) Movement(movementRequestDto *dto.MovementRequestDTO, createdAt time.Time) ([]*domain.Movement, error) {
	ctx := context.Background()

	recurrence := movementRequestDto.Recurrence
	ids := make([]string, recurrence)
	dueDates := make([]time.Time, recurrence)
	values := make([]int, recurrence)

	for i := 0; i < recurrence; i++ {
		ids[i] = uuid.New().String()
		dueDates[i] = movementRequestDto.DueDate.AddDate(0, i, 0)
		values[i] = int(movementRequestDto.Value * 100)
	}

	query := `
        INSERT INTO tb_movement 
            (pk_st_id, it_value, dt_due_date, st_type, st_account_id, st_description, dt_created_at)
        SELECT 
            unnest($1::uuid[]), 
            unnest($2::integer[]), 
            unnest($3::date[]), 
            unnest($4::text[]), 
            unnest($5::text[]), 
            unnest($6::text[]), 
            unnest($7::timestamp[])
        RETURNING *
    `

	rows, err := repository.db.Query(
		ctx,
		query,
		ids,
		values,
		dueDates,
		RepeatString(string(movementRequestDto.Type), recurrence),
		RepeatString(movementRequestDto.AccountID, recurrence),
		RepeatString(movementRequestDto.Description, recurrence),
		RepeatTime(createdAt, recurrence),
	)

	if err != nil {
		return nil, fmt.Errorf("erro ao inserir movimentos: %w", err)
	}
	defer rows.Close()

	var movements []*domain.Movement
	for rows.Next() {
		var m domain.Movement
		err := rows.Scan(
			&m.ID,
			&m.Value,
			&m.DueDate,
			&m.Type,
			&m.AccountID,
			&m.Description,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler movimento: %w", err)
		}
		movements = append(movements, &m)
	}

	return movements, nil
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
