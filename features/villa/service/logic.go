package service

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/villa"
	"github.com/go-playground/validator/v10"
)

type villaService struct {
	villaRepository villa.RepositoryInterface
	validate        *validator.Validate
}

func New(repo villa.RepositoryInterface) villa.ServiceInterface {
	return &villaService{
		villaRepository: repo,
		validate:        validator.New(),
	}
}

// Create implements villa.ServiceInterface
func (service *villaService) Create(input villa.CoreVilla) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.villaRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}
