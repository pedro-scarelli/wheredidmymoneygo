package accountrepository

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/adapter/postgres"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.AccountRepository {
	return &repository{
		db: db,
	}
}
