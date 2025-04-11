package userservice

import "github.com/pedro-scarelli/go_login/core/domain"

type service struct {
	usecase domain.UserUseCase
}

func New(usecase domain.UserUseCase) domain.UserService {
	return &service{
		usecase: usecase,
	}
}
