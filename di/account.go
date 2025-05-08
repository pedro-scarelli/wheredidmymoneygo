package di

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/http/accountservice"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres"
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres/accountrepository"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/usecase/accountusecase"
)

func ConfigAccountDI(conn postgres.PoolInterface) domain.AccountService {
	accountRepository := accountrepository.New(conn)
	accountUseCase := accountusecase.New(accountRepository)
	accountService := accountservice.New(accountUseCase)

	return accountService
}
