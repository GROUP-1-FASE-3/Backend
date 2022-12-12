package repository

import (
	"github.com/GROUP-1-FASE-3/Backend/features/creditcard"
	"gorm.io/gorm"
)

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

type User struct {
	User_Name    string
	Email        string
	Password     string
	Gender       string
	Phone_Number string
	User_Images  string
	CreditCard   CreditCard
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

func fromCore(dataCore creditcard.CoreCreditCard) CreditCard {
	creditcardGorm := CreditCard{
		Type:   dataCore.Type,
		Name:   dataCore.Name,
		Number: dataCore.Number,
		Cvv:    dataCore.Cvv,
		Month:  dataCore.Month,
		Year:   dataCore.Year,
		UserID: dataCore.User.ID,
	}
	return creditcardGorm
}
