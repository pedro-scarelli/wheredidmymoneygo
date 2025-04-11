package accountusecase

import (
	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/security"
	"math/rand"
	"time"
)

func (usecase usecase) Update(accountRequest *dto.CreateAccountRequest) (*domain.PublicAccount, error) {
	hashedPassword, err := security.HashPassword(accountRequest.Password)
	if err != nil {
		return nil, err
	}
	accountRequest.Password = hashedPassword

	account, err := usecase.repository.Create(accountRequest, rand.Intn(10000000), time.Now().UTC())

	if err != nil {
		return nil, err
	}

	return account, nil
}
