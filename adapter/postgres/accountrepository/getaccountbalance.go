package accountrepository

import (
	"context"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (repository repository) GetAccountBalance(accountID string) (int, error) {
	ctx := context.Background()

	rows, err := repository.db.Query(ctx, `
        SELECT
            pk_st_id,
            it_value
        FROM tb_movement
        WHERE st_account_id = $1
          AND dt_due_date >= date_trunc('month', current_date)
          AND dt_due_date <  date_trunc('month', current_date) + interval '1 month'
    `, accountID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var movements []domain.Movement
	var balance int

	for rows.Next() {
		var m domain.Movement
		err := rows.Scan(&m.ID, &m.Value)
		if err != nil {
			return 0, err
		}
		movements = append(movements, m)
		balance += m.Value
	}

	if err := rows.Err(); err != nil {
		return 0, err
	}

	if len(movements) == 0 {
		return  0, nil
	}

	return balance, nil
}