package di

import (
	"github.com/pedro-scarelli/go_login/adapter/http/userservice"
	"github.com/pedro-scarelli/go_login/adapter/postgres"
	"github.com/pedro-scarelli/go_login/adapter/postgres/userrepository"
	"github.com/pedro-scarelli/go_login/core/domain"
	"github.com/pedro-scarelli/go_login/core/usecase/userusecase"
)

func ConfigUserDI(conn postgres.PoolInterface) domain.UserService {
	userRepository := userrepository.New(conn)
	userUseCase := userusecase.New(userRepository)
	userService := userservice.New(userUseCase)

	return userService
}
