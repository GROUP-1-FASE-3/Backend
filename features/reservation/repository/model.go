package repository

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	Start_Date   string
	End_Date     string
	Price        uint
	Total_Price  uint
	VillaID      uint
	UserID       uint
	CreditCardID uint
	User         User
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
