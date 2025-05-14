package di

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/http/authenticationservice"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres/accountrepository"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/usecase/authenticationusecase"
)

func ConfigAuthenticationDI(conn postgres.PoolInterface) domain.AuthenticationService {
	accountRepository := accountrepository.New(conn)
	authenticationUseCase := authenticationusecase.New(accountRepository)
	authenticationService := authenticationservice.New(authenticationUseCase)

	return authenticationService
}
