package accountrepository

import (
	"context"
	"fmt"

	"dario.cat/mergo"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
)

func (repository repository) Update(updateAccountRequestDto *dto.UpdateAccountRequestDTO) (*domain.PublicAccount, error) {
	account, err := repository.GetByID(int(updateAccountRequestDto.ID))
	if err != nil {
		return nil, fmt.Errorf("conta não encontrada")
	}

	updateData := &domain.Account{
		ID:        updateAccountRequestDto.ID,
		Number:    account.Number,
		CPF:       account.CPF,
		Email:     account.Email,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		DeletedAt: account.DeletedAt,
	}

    if updateAccountRequestDto.FirstName != nil {
        updateData.FirstName = *updateAccountRequestDto.FirstName
    }
    if updateAccountRequestDto.LastName != nil {
        updateData.LastName = *updateAccountRequestDto.LastName
    }
    if updateAccountRequestDto.Password != nil {
        updateData.Password = *updateAccountRequestDto.Password
	}

	err = mergo.MergeWithOverwrite(account, updateData)
	if err != nil {
		return nil, fmt.Errorf("erro ao mapear propiedades: %v", err)
	}

	fmt.Printf("ID da conta pra ser atualizada: %v", updateAccountRequestDto.ID)
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
