package villa

import "github.com/labstack/echo/v4"

type CoreVilla struct {
	ID               uint
	Villa_Name       string
	Price            uint
	Description      string
	Address          string
	Villa_Images1    string
	Villa_Images2    string
	Villa_Images3    string
	Detail_Guest     uint
	Detail_Bedroom   uint
	Detail_Bed       uint
	Detail_Bath      uint
	Detail_Kitchen   string
	Detail_Pool      string
	Detail_Wifi      string
	Detail_Workspace string
	UserID           uint
	User             User
	Rating           CoreRating
}

type User struct {
	ID        uint
	User_Name string
}

type CoreRating struct {
	ID     uint
	Rating uint
}

type ServiceInterface interface {
	Create(input CoreVilla, c echo.Context) (err error)
	GetAll() (data []CoreVilla, err error)
	GetById(id int) (data CoreVilla, err error)
	UpdateVilla(input CoreVilla, id int) (err error)
	DeleteVilla(id int) (err error)
	GetAllByID(id int) (data []CoreVilla, err error)
}

type RepositoryInterface interface {
	Create(input CoreVilla) (row int, err error)
	GetAll() (data []CoreVilla, err error)
	GetById(id int) (data CoreVilla, err error)
	UpdateVilla(input CoreVilla, id int) (err error)
	DeleteVilla(id int) (row int, err error)
	GetAllByID(id int) (data []CoreVilla, err error)
}
