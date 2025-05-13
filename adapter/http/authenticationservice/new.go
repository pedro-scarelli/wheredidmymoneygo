package authenticationservice

import "github.com/pedro-scarelli/wheredidmymoneygo/core/domain"

type service struct {
	usecase domain.AuthenticationUseCase
}

func New(usecase domain.AuthenticationUseCase) domain.AuthenticationService {
	return &service{
		usecase: usecase,
	}
}
