package userusecase

import "github.com/pedro-scarelli/go_login/core/domain"

type usecase struct {
	repository domain.UserRepository
}

func New(repository domain.UserRepository) domain.UserUseCase {
	return &usecase{
		repository: repository,
	}
}
