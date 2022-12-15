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
	User        User
	Reservation []Reservation
}

type User struct {
	gorm.Model
	User_Name    string
	Email        string `gorm:"type:varchar(200)"`
	Password     string
	Gender       string
	Phone_Number string
	User_Images  string
	CreditCard   []CreditCard
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

// mengubah struct model gorm ke struct core
func (dataModel *CreditCard) toCore() creditcard.CoreCreditCard {
	return creditcard.CoreCreditCard{
		ID:     dataModel.ID,
		Type:   dataModel.Type,
		Name:   dataModel.Name,
		Number: dataModel.Number,
		Cvv:    dataModel.Cvv,
		Month:  dataModel.Month,
		Year:   dataModel.Year,
		User:   dataModel.User.toCore2(),
	}
}

// to Core User
func (dataModel2 *User) toCore2() creditcard.CoreUser {
	return creditcard.CoreUser{
		ID:        dataModel2.ID,
		User_Name: dataModel2.User_Name,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []CreditCard) []creditcard.CoreCreditCard {
	var dataCore []creditcard.CoreCreditCard
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

// append to CoreUser
func toCoreList2(dataModel2 []User) []creditcard.CoreUser {
	var dataCore2 []creditcard.CoreUser
	for _, v := range dataModel2 {
		dataCore2 = append(dataCore2, v.toCore2())
	}
	return dataCore2
}
