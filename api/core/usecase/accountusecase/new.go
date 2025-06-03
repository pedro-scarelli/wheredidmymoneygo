package accountusecase

import "github.com/pedro-scarelli/wheredidmymoneygo/core/domain"

type usecase struct {
	repository domain.AccountRepository
}

func New(repository domain.AccountRepository) domain.AccountUseCase {
	return &usecase{
		repository: repository,
	}
}
