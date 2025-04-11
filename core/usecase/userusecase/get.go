package userusecase

import (
	"github.com/pedro-scarelli/go_login/core/domain"
	"github.com/pedro-scarelli/go_login/core/dto"
)

func (usecase usecase) Get(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination[[]domain.User], error) {
	users, err := usecase.repository.Get(paginationRequest)

	if err != nil {
		return nil, err
	}

	return users, nil
}
