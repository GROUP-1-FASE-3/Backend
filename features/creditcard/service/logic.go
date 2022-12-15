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

// Post
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

// Get by (ID)
func (service *creditcardService) GetById(id int) (data creditcard.CoreCreditCard, err error) {
	data, errGet := service.creditcardRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get creditcard by id data, error query")
	}
	return data, nil
}

// Update
func (service *creditcardService) UpdateCreditCard(dataCore creditcard.CoreCreditCard, id int) (err error) {
	errUpdate := service.creditcardRepository.UpdateCreditCard(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

// Delete
func (service *creditcardService) DeleteCreditCard(id int) (err error) {
	_, errDel := service.creditcardRepository.DeleteCreditCard(id)
	if errDel != nil {
		return errors.New("failed delete creditcard, error query")
	}
	return nil
}

func (service *creditcardService) GetByUserId(id int) (data []creditcard.CoreCreditCard, err error) {
	data, errGet := service.creditcardRepository.GetByUserId(id)
	if errGet != nil {
		return data, errors.New("failed get creditcard by id data, error query")
	}
	return data, nil
}
