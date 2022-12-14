package user

import "github.com/labstack/echo/v4"

type UserCore struct {
	ID           uint
	User_Name    string `validate:"required"`
	Email        string `validate:"required,email"`
	Password     string `validate:"required"`
	Gender       string
	Phone_Number string
	User_Images  string
}

type RepositoryInterface interface {
	Create(input UserCore) (row int, err error)
	Delete(data UserCore, id int) (row int, err error)
	GetByID(id int) (data UserCore, err error)
	Update(id int, data UserCore) (row int, err error)
}

type ServiceInterface interface {
	Create(input UserCore) (err error)
	Delete(data UserCore, id int) (err error)
	GetByID(id int) (data UserCore, err error)
	Update(id int, data UserCore, c echo.Context) (err error)
}
