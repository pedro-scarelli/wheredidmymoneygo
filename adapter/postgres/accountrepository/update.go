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
	ctx := context.Background()

	if err != nil {
		return nil, fmt.Errorf("Conta não encontrada")
	}

	err = mergo.MergeWithOverwrite(&account, updateAccountRequest)

	if err != nil {
		return nil, fmt.Errorf("Erro ao mapear propiedades")
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
