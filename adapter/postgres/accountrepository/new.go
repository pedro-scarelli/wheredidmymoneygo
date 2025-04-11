package userrepository

import (
	"github.com/pedro-scarelli/go_login/adapter/postgres"
	"github.com/pedro-scarelli/go_login/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.UserRepository {
	return &repository{
		db: db,
	}
}
