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

// GetAll implements villa.ServiceInterface
func (service *villaService) GetAll() (data []villa.CoreVilla, err error) {
	data, err = service.villaRepository.GetAll()
	return

}

// Get by ID
func (service *villaService) GetById(id int) (data villa.CoreVilla, err error) {
	data, errGet := service.villaRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get user by id data, error query")
	}
	return data, nil
}

// Update
func (service *villaService) UpdateVilla(dataCore villa.CoreVilla, id int) (err error) {
	errUpdate := service.villaRepository.UpdateVilla(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}
