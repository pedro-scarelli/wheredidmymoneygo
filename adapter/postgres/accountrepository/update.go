package userrepository

import (
	"context"
	"github.com/pedro-scarelli/go_login/core/domain"
	"github.com/pedro-scarelli/go_login/core/dto"
	"time"
)

func (repository repository) Update(userRequest *dto.CreateUserRequest) (*domain.PublicUser, error) {
	ctx := context.Background()
	user := domain.User{}

	err := repository.db.QueryRow(
		ctx,
		`INSERT INTO tb_user
		(st_first_name, st_last_name, st_cpf, st_email, st_password, it_number, db_balance, dt_created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		returning *`,
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.CPF,
		userRequest.Email,
		userRequest.Password,
		1341,
		0,
		time.Now().UTC(),
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.CPF,
		&user.Email,
		&user.Password,
		&user.Number,
		&user.Balance,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	publicUser := domain.PublicUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CPF:       user.CPF,
		Email:     user.Email,
		Number:    user.Number,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}

	return &publicUser, nil
}
