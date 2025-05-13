package accountrepository

import (
	"context"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
)

func (repository repository) Get(paginationRequestDTO *dto.PaginationRequestParams) (*domain.Pagination[[]domain.PublicAccount], error) {
	ctx := context.Background()
	page := paginationRequestDTO.Page
	itemsPerPage := paginationRequestDTO.ItemsPerPage

	offset := (page - 1) * itemsPerPage

	query := `
        SELECT
            pk_it_id, st_first_name, st_last_name, st_cpf, st_email,
            it_number, it_balance, dt_created_at,
            COUNT(*) OVER() AS total_count
        FROM tb_account
        ORDER BY pk_it_id
        LIMIT $1 OFFSET $2
    `
	rows, err := repository.db.Query(ctx, query, itemsPerPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []domain.PublicAccount
	var totalCount int32
	for rows.Next() {
		var acc domain.PublicAccount
		err := rows.Scan(
			&acc.ID,
			&acc.FirstName,
			&acc.LastName,
			&acc.CPF,
			&acc.Email,
			&acc.Number,
			&acc.Balance,
			&acc.CreatedAt,
			&totalCount,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	return &domain.Pagination[[]domain.PublicAccount]{
		TotalItems: totalCount,
		Page:       int32(page),
		Data:       accounts,
	}, nil
}
