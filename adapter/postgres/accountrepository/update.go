package accountrepository

import (
	"context"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
	"time"
)

func (repository repository) Update(accountRequest *dto.CreateAccountRequest) (*domain.PublicAccount, error) {
	ctx := context.Background()
	account := domain.Account{}

	err := repository.db.QueryRow(
		ctx,
		`INSERT INTO tb_account
		(st_first_name, st_last_name, st_cpf, st_email, st_password, it_number, db_balance, dt_created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		returning *`,
		accountRequest.FirstName,
		accountRequest.LastName,
		accountRequest.CPF,
		accountRequest.Email,
		accountRequest.Password,
		1341,
		0,
		time.Now().UTC(),
	).Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.CPF,
		&account.Email,
		&account.Password,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	publicAccount := domain.PublicAccount{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		CPF:       account.CPF,
		Email:     account.Email,
		Number:    account.Number,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}

	return &publicAccount, nil
}
