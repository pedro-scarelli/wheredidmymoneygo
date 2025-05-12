package accountrepository

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.AccountRepository {
	return &repository{
		db: db,
	}
}
