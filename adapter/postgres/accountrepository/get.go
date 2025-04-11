package userrepository

import (
	"context"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/pedro-scarelli/go_login/core/domain"
	"github.com/pedro-scarelli/go_login/core/dto"
)

func (repository repository) Get(pagination *dto.PaginationRequestParms) (*domain.Pagination[[]domain.User], error) {
	ctx := context.Background()
	users := []domain.User{}
	total := int32(0)
	pagin := paginate.Instance(pagination)
	query, queryCount := pagin.
		Query(`
		SELECT pk_it_id, st_first_name, st_last_name, st_cpf, st_email, it_number, db_balance, dt_created_at 
		FROM tb_user`,
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
			user := domain.User{}

			rows.Scan(
				&user.ID,
				&user.FirstName,
				&user.LastName,
				&user.CPF,
				&user.Email,
				&user.Number,
				&user.Balance,
				&user.CreatedAt,
			)

			users = append(users, user)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.User]{
		Items: users,
		Total: total,
	}, nil
}
