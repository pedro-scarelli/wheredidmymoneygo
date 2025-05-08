package accountrepository

import (
	"context"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
)

func (repository repository) Get(pagination *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Account], error) {
	ctx := context.Background()
	accounts := []domain.Account{}
	total := int32(0)
	pagin := paginate.Instance(pagination)
	query, queryCount := pagin.
		Query(`
		SELECT pk_it_id, st_first_name, st_last_name, st_cpf, st_email, it_number, db_balance, dt_created_at 
		FROM tb_account`,
		).
		Sort(pagination.Sort).
		Desc(pagination.Descending).
		Page(pagination.Page).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "name", "description").
		Select()

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			account := domain.Account{}

			rows.Scan(
				&account.ID,
				&account.FirstName,
				&account.LastName,
				&account.CPF,
				&account.Email,
				&account.Number,
				&account.Balance,
				&account.CreatedAt,
			)

			accounts = append(accounts, account)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.Account]{
		Items: accounts,
		Total: total,
	}, nil
}
