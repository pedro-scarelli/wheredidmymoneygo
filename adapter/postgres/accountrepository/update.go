package accountrepository

import (
	"context"
	"fmt"
	"strings"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
)

func (r repository) Update(request *dto.UpdateAccountRequest) (*domain.PublicAccount, error) {
	ctx := context.Background()
	account := domain.Account{}

	fields := map[string]*string{
		"st_first_name": request.FirstName,
		"st_last_name":  request.LastName,
		"st_password":   request.Password,
	}

	setClauses := make([]string, 0, len(fields)+1)
	args := make([]any, 0, len(fields)+2)
	idx := 1

	for col, ptr := range fields {
		if ptr != nil && *ptr != "" {
			setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, idx))
			args = append(args, *ptr)
			idx++
		}
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf(
		"UPDATE tb_account SET %s WHERE pk_it_id = $%d RETURNING pk_it_id, st_first_name, st_last_name, st_cpf, st_email, st_password, it_number, it_balance, dt_created_at",
		strings.Join(setClauses, ", "),
		idx,
	)
	args = append(args, request.ID)

	err := r.db.QueryRow(ctx, query, args...).Scan(
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
