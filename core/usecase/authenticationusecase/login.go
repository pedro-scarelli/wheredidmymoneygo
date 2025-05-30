package authenticationusecase

import (
	"fmt"

	requestDto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/authentication/request"
	responseDto "github.com/pedro-scarelli/wheredidmymoneygo/core/dto/authentication/response"
	"github.com/pedro-scarelli/wheredidmymoneygo/core/security"
)

func (usecase usecase) Login(loginRequestDto *requestDto.LoginRequestDTO) (*responseDto.LoginResponseDTO, error) {
	account, err := usecase.repository.GetAccountByEmail(loginRequestDto.Email)
	if err != nil {
		return nil, fmt.Errorf("incorrect email or password")
	}

	if security.IsPasswordIncorrect(loginRequestDto.Password, account.Password) {
		return nil, fmt.Errorf("incorrect email or password")
	}

	jwtToken := security.GenerateJwtToken(account.ID)

	loginResponseDto := responseDto.LoginResponseDTO{Token: jwtToken}

	return &loginResponseDto, nil
}

