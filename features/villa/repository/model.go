package repository

import (
	"github.com/GROUP-1-FASE-3/Backend/features/villa"
	"gorm.io/gorm"
)

type Villa struct {
	gorm.Model
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
	Ratings          []Rating
	Reservations     []Reservation
	Rating           Rating
	User             User
}

type User struct {
	gorm.Model
	User_Name    string
	Email        string `gorm:"type:varchar(200)"`
	Password     string
	Gender       string
	Phone_Number string
	User_Images  string
	Villas       []Villa
}

type Rating struct {
	gorm.Model
	Rating  uint
	Comment string
	VillaID uint
	UserID  uint
}

type Reservation struct {
	gorm.Model
	Start_Date   string
	End_Date     string
	Price        uint
	Total_Price  uint
	VillaID      uint
	UserID       uint
	CreditCardID uint
}

// mengubah struct core ke struct model gorm
func fromCore(dataCore villa.CoreVilla) Villa {
	villaGorm := Villa{
		Villa_Name:       dataCore.Villa_Name,
		Price:            dataCore.Price,
		Description:      dataCore.Description,
		Address:          dataCore.Address,
		Villa_Images1:    dataCore.Villa_Images1,
		Villa_Images2:    dataCore.Villa_Images2,
		Villa_Images3:    dataCore.Villa_Images3,
		Detail_Guest:     dataCore.Detail_Guest,
		Detail_Bedroom:   dataCore.Detail_Bedroom,
		Detail_Bed:       dataCore.Detail_Bed,
		Detail_Bath:      dataCore.Detail_Bath,
		Detail_Kitchen:   dataCore.Detail_Kitchen,
		Detail_Pool:      dataCore.Detail_Pool,
		Detail_Wifi:      dataCore.Detail_Wifi,
		Detail_Workspace: dataCore.Detail_Workspace,
		UserID:           dataCore.UserID,
	}
	return villaGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Villa) toCore() villa.CoreVilla {
	return villa.CoreVilla{
		ID:               dataModel.ID,
		Villa_Name:       dataModel.Villa_Name,
		Price:            dataModel.Price,
		Description:      dataModel.Description,
		Address:          dataModel.Address,
		Villa_Images1:    dataModel.Villa_Images1,
		Villa_Images2:    dataModel.Villa_Images2,
		Villa_Images3:    dataModel.Villa_Images3,
		Detail_Guest:     dataModel.Detail_Guest,
		Detail_Bedroom:   dataModel.Detail_Bedroom,
		Detail_Bed:       dataModel.Detail_Bed,
		Detail_Bath:      dataModel.Detail_Bath,
		Detail_Kitchen:   dataModel.Detail_Kitchen,
		Detail_Pool:      dataModel.Detail_Pool,
		Detail_Wifi:      dataModel.Detail_Wifi,
		Detail_Workspace: dataModel.Detail_Workspace,
		UserID:           dataModel.UserID,
		Rating: villa.CoreRating{
			ID:     dataModel.Rating.ID,
			Rating: dataModel.Rating.Rating,
		},
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Villa) []villa.CoreVilla {
	var dataCore []villa.CoreVilla
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
