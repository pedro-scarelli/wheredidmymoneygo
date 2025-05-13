package accountrepository

import (
	"context"
	"fmt"

	"dario.cat/mergo"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
)

func (repository repository) Update(updateAccountRequest *dto.UpdateAccountRequest) (*domain.PublicAccount, error) {
	account, err := repository.GetByID(int(updateAccountRequest.ID))
	if err != nil {
		return nil, fmt.Errorf("Conta não encontrada")
	}

	updateData := &domain.Account{
		ID:        updateAccountRequest.ID,
		FirstName: updateAccountRequest.FirstName,
		LastName:  updateAccountRequest.LastName,
		Number:    account.Number,
		CPF:       account.CPF,
		Email:     account.Email,
		Password:  updateAccountRequest.Password,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		DeletedAt: account.DeletedAt,
	}

	err = mergo.MergeWithOverwrite(account, updateData)
	if err != nil {
		return nil, fmt.Errorf("Erro ao mapear propiedades: %v", err)
	}

	fmt.Printf("ID da conta pra ser atualizada: %v", updateAccountRequest.ID)
	ctx := context.Background()
	query := `
        UPDATE tb_account
        SET st_first_name = $1,
            st_last_name  = $2,
            st_cpf        = $3,
            st_email      = $4,
			st_password   = $5,
            it_number     = $6,
            it_balance    = $7,
        WHERE pk_it_id = $8;
    `
	err = repository.db.QueryRow(
		ctx,
		query,
		account.FirstName,
		account.LastName,
		account.CPF,
		account.Email,
		account.Password,
		account.Number,
		account.Balance,
		account.CreatedAt,
		account.ID,
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
