package repository

import (
	"time"

	"github.com/GROUP-1-FASE-3/Backend/features/reservation"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	Start_Date   time.Time
	End_Date     time.Time
	VillaPrice   uint
	Jumlah_malam uint
	Total_Price  uint
	VillaID      uint
	Villa        Villa
	UserID       uint
	User         User
	CreditCardID uint
	CreditCard   CreditCard
}

type User struct {
	gorm.Model
	User_Name    string
	Email        string
	Password     string
	Gender       string
	Phone_Number string
	User_Images  string
	Reservations []Reservation
}

type Villa struct {
	gorm.Model
	Villa_Name     string
	Price          uint
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
	Reservations   []Reservation
}

type CreditCard struct {
	gorm.Model
	Type        string
	Name        string
	Number      string
	Cvv         uint
	Month       uint
	Year        uint
	UserID      uint
	Reservation []Reservation
}

func ReserveCoreToModel(data reservation.ReservationCore) Reservation {
	reserveData := Reservation{
		Start_Date:   data.Start_Date,
		End_Date:     data.End_Date,
		VillaPrice:   data.Price,
		Total_Price:  data.Total_price,
		VillaID:      data.VillaID,
		UserID:       data.UserID,
		CreditCardID: data.CreditCardID,
	}

	return reserveData
}

func (dataModel *Reservation) toCore() reservation.ReservationCore {
	return reservation.ReservationCore{
		UserID:      dataModel.UserID,
		ID:          dataModel.ID,
		Villa:       dataModel.Villa.toCoreV(),
		Total_price: dataModel.Total_Price,
		Start_Date:  dataModel.Start_Date,
		End_Date:    dataModel.End_Date,
	}
}

func (dataModel *Villa) toCoreV() reservation.VillaCore {
	return reservation.VillaCore{
		Villa_Name: dataModel.Villa_Name,
		Price:      dataModel.Price,
	}
}

func toCoreList(dataModel []Reservation) []reservation.ReservationCore {
	var dataCore []reservation.ReservationCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
