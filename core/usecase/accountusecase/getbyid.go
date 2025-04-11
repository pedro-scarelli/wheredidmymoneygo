package userusecase

import (
	"github.com/pedro-scarelli/go_login/core/domain"
)

func (usecase usecase) GetByID(userID int) (*domain.User, error) {
	user, err := usecase.repository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
