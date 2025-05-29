package accountrepository

import (
	"context"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

func (repository repository) GetAccountByEmail(email string) (*domain.Account, error) {
	ctx := context.Background()
	row := repository.db.QueryRow(ctx,
		`select 
		pk_st_id, st_first_name, st_last_name, st_cpf, st_email, it_number, dt_created_at, st_password
		from tb_account
		where st_email = $1`,
		email)

	return scanIntoAccount(row)
}
