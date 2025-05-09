package accountrepository

import (
	"context"
	"fmt"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
	"strconv"
)

func (repository repository) Get(paginationRequestDTO *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Account], error) {
	ctx := context.Background()
	page := paginationRequestDTO.Page
	itemsPerPage := paginationRequestDTO.ItemsPerPage

	offset := (page - 1) * itemsPerPage

	fmt.Printf("page: %v\n", strconv.Itoa(page))
	fmt.Printf("items: %v\n", strconv.Itoa(itemsPerPage))
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

	var accounts []domain.Account
	var totalCount int32
	for rows.Next() {
		var acc domain.Account
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

	return &domain.Pagination[[]domain.Account]{
		Total:      totalCount,
		TotalItems: int32(len(accounts)),
		Page:       int32(page),
		Data:       accounts,
	}, nil
}
