package accountusecase

import (
	"math/rand"
	"time"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/domain"
	dto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/account/request"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/security"
)

func (usecase usecase) Create(createAccountRequestDto *dto.CreateAccountRequestDTO) (*domain.PublicAccount, error) {
	hashedPassword, err := security.HashPassword(createAccountRequestDto.Password)
	if err != nil {
		return nil, err
	}
	createAccountRequestDto.Password = hashedPassword

	publicAccount, err := usecase.repository.Create(createAccountRequestDto, rand.Intn(10000000), time.Now().UTC())

	if err != nil {
		return nil, err
	}

	return publicAccount, nil
}
