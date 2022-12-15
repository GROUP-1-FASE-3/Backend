package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/GROUP-1-FASE-3/Backend/features/villa"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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
func (service *villaService) Create(input villa.CoreVilla, c echo.Context) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	var Villa_Images1 string = "villa_images1"
	var Villa_Images2 string = "villa_images2"
	var Villa_Images3 string = "villa_images3"
	//Upload File
	file, _ := c.FormFile(Villa_Images1)
	if file != nil {
		res, err := helper.UploadVilla(c, Villa_Images1)
		if err != nil {
			return errors.New("Registration Failed. Cannot Upload Data Images1.")
		}
		log.Print(res)
		input.Villa_Images1 = res
	} else {
		input.Villa_Images1 = "https://fmz-airbnb-bucket.s3.ap-southeast-1.amazonaws.com/dummy.png"
	} // end upload file

	file2, _ := c.FormFile(Villa_Images2)
	if file2 != nil {
		res, err := helper.UploadVilla(c, Villa_Images2)
		if err != nil {
			return errors.New("Registration Failed. Cannot Upload Data Images2.")
		}
		log.Print(res)
		input.Villa_Images2 = res
	} else {
		input.Villa_Images2 = "https://fmz-airbnb-bucket.s3.ap-southeast-1.amazonaws.com/dummy.png"
	} // end upload file

	file3, _ := c.FormFile(Villa_Images3)
	if file3 != nil {
		res, err := helper.UploadVilla(c, Villa_Images3)
		if err != nil {
			return errors.New("Registration Failed. Cannot Upload Data Images3.")
		}
		log.Print(res)
		input.Villa_Images3 = res
	} else {
		input.Villa_Images3 = "https://fmz-airbnb-bucket.s3.ap-southeast-1.amazonaws.com/dummy.png"
	} // end upload file

	fmt.Println("hasil input create", input)
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
		return data, errors.New("failed get villa by id data, error query")
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

// Delete
func (service *villaService) DeleteVilla(id int) (err error) {
	_, errDel := service.villaRepository.DeleteVilla(id)
	if errDel != nil {
		return errors.New("failed delete villa, error query")
	}
	return nil
}

func (service *villaService) GetAllByID(user_id int) (data []villa.CoreVilla, err error) {
	data, err = service.villaRepository.GetAllByID(user_id)
	return

}
