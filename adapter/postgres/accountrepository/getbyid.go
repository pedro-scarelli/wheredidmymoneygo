package userrepository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pedro-scarelli/go_login/core/domain"
)

func (repository repository) GetByID(userID int) (*domain.User, error) {
	ctx := context.Background()
	row := repository.db.QueryRow(ctx,
		`select 
		pk_it_id, st_first_name, st_last_name, st_cpf, st_email, it_number, db_balance, dt_created_at
		from tb_user
		where pk_it_id = $1`,
		userID)

	return scanIntoUser(row)
}

func scanIntoUser(row pgx.Row) (*domain.User, error) {
	user := new(domain.User)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.CPF,
		&user.Email,
		&user.Number,
		&user.Balance,
		&user.CreatedAt)

	return user, err
}
