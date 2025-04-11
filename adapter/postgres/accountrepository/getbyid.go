package accountrepository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (repository repository) GetByID(accountID int) (*domain.Account, error) {
	ctx := context.Background()
	row := repository.db.QueryRow(ctx,
		`select 
		pk_it_id, st_first_name, st_last_name, st_cpf, st_email, it_number, db_balance, dt_created_at
		from tb_account
		where pk_it_id = $1`,
		accountID)

	return scanIntoAccount(row)
}

func scanIntoAccount(row pgx.Row) (*domain.Account, error) {
	account := new(domain.Account)
	err := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.CPF,
		&account.Email,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
