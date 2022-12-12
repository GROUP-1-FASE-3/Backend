package service

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/auth"
	"github.com/GROUP-1-FASE-3/Backend/middlewares"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"
)

type authService struct {
	authData auth.RepositoryInterface
}

func New(data auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authData: data,
	}
}

func (service *authService) Login(email, password string) (loginData auth.Core, err error) {
	if email == "" || password == "" {
		return loginData, errors.New("email dan password harus diisi")
	}

	result, err := service.authData.FindUser(email, password)
	if err != nil {
		return loginData, err
	}

	cekPass := helper.CheckPasswordHash(password, result.Password)

	if !cekPass {
		return loginData, errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(result.ID))
	if errToken != nil {
		return loginData, errToken
	}

	var DataLogin auth.Core
	DataLogin.ID = result.ID
	DataLogin.User_name = result.User_name
	DataLogin.Email = result.Email
	DataLogin.Token = token

	return DataLogin, err
}
