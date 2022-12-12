package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	User_Name    string
	Email        string
	Password     string
	Gender       string
	Phone_Number string
	User_Images  string
	CreditCard   CreditCard
	Reservations []Reservation
	Villas       []Villa
	Ratings      []Rating
}

type CreditCard struct {
	Type   string
	Name   string
	Number string
	Cvv    string
	Month  string
	Year   string
	UserID uint
}

type Reservation struct {
	gorm.Model
	Start_Date  string
	End_Date    string
	Price       string
	Total_Price string
	VillaID     string
	UserID      string
	User        User
}

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
	UserID         string
	User           User
}

type Rating struct {
	gorm.Model
	VillaID uint
	UserID  uint
	Rating  string
	Comment string
}
