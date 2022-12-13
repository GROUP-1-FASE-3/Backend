package repository

import (
	"github.com/GROUP-1-FASE-3/Backend/features/villa"
	"gorm.io/gorm"
)

type Villa struct {
	gorm.Model
	Villa_Name     string
	Price          string
	Description    string
	Address        string
	Villa_Images1  string
	Villa_Images2  string
	Villa_Images3  string
	Detail_Guest   uint
	Detail_Bedroom uint
	Detail_Bed     uint
	Detail_Bath    uint
	Detail_Kitchen string
	Detail_Pool    string
	Detail_Wifi    string
	UserID         uint
	Ratings        []Rating
	Reservations   []Reservation
}

type User struct {
	gorm.Model
	User_Name    string
	Email        string
	Password     string
	Gender       string
	Phone_Number string
	User_Images  string
	Villas       []Villa
}

type Rating struct {
	gorm.Model
	VillaID uint
	UserID  uint
	Rating  string
	Comment string
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
		Villa_Name:     dataCore.Villa_Name,
		Price:          dataCore.Price,
		Description:    dataCore.Description,
		Address:        dataCore.Address,
		Villa_Images1:  dataCore.Villa_Images1,
		Villa_Images2:  dataCore.Villa_Images2,
		Villa_Images3:  dataCore.Villa_Images3,
		Detail_Guest:   dataCore.Detail_Guest,
		Detail_Bedroom: dataCore.Detail_Bedroom,
		Detail_Bed:     dataCore.Detail_Bed,
		Detail_Bath:    dataCore.Detail_Bath,
		Detail_Kitchen: dataCore.Detail_Kitchen,
		Detail_Pool:    dataCore.Detail_Pool,
		Detail_Wifi:    dataCore.Detail_Wifi,
		UserID:         dataCore.User.ID,
	}
	return villaGorm
}
