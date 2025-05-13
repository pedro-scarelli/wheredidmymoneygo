package domain

import (
	"net/http"

	"github.com/pedro-scarelli/wheredidmymoneygo/core/dto"
)

type AuthenticationtService interface {
	Login(response http.ResponseWriter, request *http.Request)
}

type AuthenticationtUseCase interface {
	Login(accountRequest *dto.LoginRequestDTO) (*LoginResponseDTO, error)
}

type AuthenticationtRepository interface {
	Login(accountRequest *dto.LoginRequestDTO) (*PublicAccount, error)
}
