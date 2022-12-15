package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/creditcard"

type CreditCardResponse struct {
	ID     uint         `json:"id"`
	Type   string       `json:"type"`
	Name   string       `json:"name"`
	Number string       `json:"number"`
	Cvv    uint         `json:"cvv"`
	Month  uint         `json:"month"`
	Year   uint         `json:"year"`
	Users  UserResponse `json:"users"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	User_Name string `json:"user_name"`
}

// Credit Card from Core
func fromCore(dataCore creditcard.CoreCreditCard) CreditCardResponse {
	return CreditCardResponse{
		ID:     dataCore.ID,
		Type:   dataCore.Type,
		Name:   dataCore.Name,
		Number: dataCore.Number,
		Cvv:    dataCore.Cvv,
		Month:  dataCore.Month,
		Year:   dataCore.Year,
		Users:  fromCore2(dataCore.User),
	}
}

func fromCoreList(dataCore []creditcard.CoreCreditCard) []CreditCardResponse {
	var dataResponse []CreditCardResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

// User from Core
func fromCore2(dataCore2 creditcard.CoreUser) UserResponse {
	return UserResponse{
		ID:        dataCore2.ID,
		User_Name: dataCore2.User_Name,
	}
}

func fromCoreList2(dataCore2 []creditcard.CoreUser) []UserResponse {
	var dataResponse2 []UserResponse
	for _, v := range dataCore2 {
		dataResponse2 = append(dataResponse2, fromCore2(v))
	}
	return dataResponse2
}
