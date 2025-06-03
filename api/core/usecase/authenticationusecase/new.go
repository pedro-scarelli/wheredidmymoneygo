package authenticationusecase

import "github.com/pedro-scarelli/wheredidmymoneygo/core/domain"

type usecase struct {
	repository domain.AccountRepository
}

func New(repository domain.AccountRepository) domain.AuthenticationUseCase {
	return &usecase{
		repository: repository,
	}
}
