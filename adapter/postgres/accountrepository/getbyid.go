package accountrepository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (repository repository) GetAccountByID(accountID string) (*domain.Account, error) {
	ctx := context.Background()
	row := repository.db.QueryRow(ctx,
		`select 
		pk_st_id, st_first_name, st_last_name, st_cpf, st_email, it_number, it_balance, dt_created_at, st_password
		from tb_account
		where pk_st_id = $1`,
		accountID)

	account, err := scanIntoAccount(row)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrAccountNotFound
	}

	return account, err
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
		&account.CreatedAt,
		&account.Password,
	)

	return account, err
}
