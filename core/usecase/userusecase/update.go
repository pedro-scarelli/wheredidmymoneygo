package userusecase

import (
	"github.com/pedro-scarelli/go_login/core/domain"
	"github.com/pedro-scarelli/go_login/core/dto"
	"github.com/pedro-scarelli/go_login/core/security"
	"math/rand"
	"time"
)

func (usecase usecase) Update(userRequest *dto.CreateUserRequest) (*domain.PublicUser, error) {
	hashedPassword, err := security.HashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}
	userRequest.Password = hashedPassword

	user, err := usecase.repository.Create(userRequest, rand.Intn(10000000), time.Now().UTC())

	if err != nil {
		return nil, err
	}

	return user, nil
}
