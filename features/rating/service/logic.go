package service

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/rating"
	"github.com/go-playground/validator/v10"
)

type ratingService struct {
	ratingRepository rating.RepositoryInterface
	validate         *validator.Validate
}

func New(repo rating.RepositoryInterface) rating.ServiceInterface {
	return &ratingService{
		ratingRepository: repo,
		validate:         validator.New(),
	}
}

// Create implements rating.ServiceInterface
func (service *ratingService) Create(input rating.CoreRating) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.ratingRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements user.ServiceInterface
func (service *ratingService) GetAll() (data []rating.CoreRating, err error) {
	data, err = service.ratingRepository.GetAll()
	return

}

// Get by ID
func (service *ratingService) GetById(id int) (data rating.CoreRating, err error) {
	data, errGet := service.ratingRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get rating by id data, error query")
	}
	return data, nil
}

// Update by id
func (service *ratingService) UpdateRating(dataCore rating.CoreRating, id int) (err error) {
	errUpdate := service.ratingRepository.UpdateRating(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

// Delete
func (service *ratingService) DeleteRating(id int) (err error) {
	_, errDel := service.ratingRepository.DeleteRating(id)
	if errDel != nil {
		return errors.New("failed delete rating, error query")
	}
	return nil
}
