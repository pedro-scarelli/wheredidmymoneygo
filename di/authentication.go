package di

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/http/authenticationservice"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres/authenticationrepository"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/usecase/authenticationusecase"
)

func ConfigAuthenticationDI(conn postgres.PoolInterface) domain.AuthenticationService {
	authenticationRepository := authenticationrepository.New(conn)
	authenticationUseCase := authenticationusecase.New(authenticationRepository)
	authenticationService := authenticationservice.New(authenticationUseCase)

	return authenticationService
}
