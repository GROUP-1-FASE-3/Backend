package service

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/user"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

func (s *userService) Create(data user.UserCore) (err error) {
	if errValidate := s.validate.Struct(data); errValidate != nil {
		return errValidate
	}

	data.Password, err = helper.HashPassword(data.Password)
	if err != nil {
		return errors.New("failed to hash")
	}

	_, errCreate := s.userRepository.Create(data)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (s *userService) Delete(data user.UserCore, id int) (err error) {
	_, errDelete := s.userRepository.Delete(data, id)
	if errDelete != nil {
		return errors.New("error delete")
	}

	return nil
}

func (s *userService) GetByID(id int) (data user.UserCore, err error) {
	dataCore, errData := s.userRepository.GetByID(id)
	if errData != nil {
		return user.UserCore{}, errData
	}
	return dataCore, nil
}
