package service

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/creditcard"
	"github.com/go-playground/validator/v10"
)

type creditcardService struct {
	creditcardRepository creditcard.RepositoryInterface
	validate             *validator.Validate
}

func New(repo creditcard.RepositoryInterface) creditcard.ServiceInterface {
	return &creditcardService{
		creditcardRepository: repo,
		validate:             validator.New(),
	}
}

func (service *creditcardService) Create(input creditcard.CoreCreditCard) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.creditcardRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}
