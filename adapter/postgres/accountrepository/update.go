package accountrepository

import (
	"context"
	"fmt"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

func (repository repository) Update(updateAccountRequestDto *dto.UpdateAccountRequestDTO) (*domain.PublicAccount, error) {
	account := domain.Account{}
	ctx := context.Background()
	query := `
		UPDATE tb_account
		SET st_first_name = $1,
		    st_last_name  = $2,
		    st_password   = $3
		WHERE pk_st_id = $4
		RETURNING pk_st_id, st_first_name, st_last_name, st_password;
	    `
	err := repository.db.QueryRow(
		ctx,
		query,
		updateAccountRequestDto.FirstName,
		updateAccountRequestDto.LastName,
		updateAccountRequestDto.Password,
		updateAccountRequestDto.ID,
	).Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao atualizar conta no banco: %w", err)
	}

	return &domain.PublicAccount{
		ID:        account.ID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		CPF:       account.CPF,
		Email:     account.Email,
		Number:    account.Number,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}, nil
}
