package authenticationusecase

import "github.com/pedro-scarelli/wheredidmymoneygo/core/domain"

type usecase struct {
	repository domain.AuthenticationRepository
}

func New(repository domain.AuthenticationRepository) domain.AuthenticationUseCase {
	return &usecase{
		repository: repository,
	}
}
