package service

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/user"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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

func (s *userService) Update(id int, input user.UserCore, c echo.Context) (err error) {
	if input.Password != "" {
		input.Password, err = helper.HashPassword(input.Password)
		if err != nil {
			return errors.New("error hashing")
		}
	}

	var User_images string = "user_images"
	file, _ := c.FormFile(User_images)
	if file != nil {
		res, err := helper.UploadVilla(c, User_images)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}

		input.User_Images = res
	} else {
		input.User_Images = "https://fmz-airbnb-bucket.s3.ap-southeast-1.amazonaws.com/dummy.png"
	}
	_, errUpdate := s.userRepository.Update(id, input)
	if errUpdate != nil {
		return errors.New("failed to insert data, error query")
	}
	return err
}
