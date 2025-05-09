package accountusecase

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
)

func (usecase usecase) Get(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Account], error) {
	accounts, err := usecase.repository.Get(paginationRequest)

	if err != nil {
		return nil, err
	}

	return accounts, nil
}
