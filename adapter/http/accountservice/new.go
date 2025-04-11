package accountservice

import "github.com/pedro-scarelli/wheredidmymoneygo/core/domain"

type service struct {
	usecase domain.AccountUseCase
}

func New(usecase domain.AccountUseCase) domain.AccountService {
	return &service{
		usecase: usecase,
	}
}
