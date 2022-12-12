package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/creditcard"

type CreditCardRequest struct {
	Type   string `json:"type" form:"type"`
	Name   string `json:"name" form:"name"`
	Number string `json:"number" form:"number"`
	Cvv    uint   `json:"cvv" form:"cvv"`
	Month  uint   `json:"month" form:"month"`
	Year   uint   `json:"year" form:"year"`
	UserID uint   `json:"user_id" form:"user_id"`
}

func toCore(data CreditCardRequest) creditcard.CoreCreditCard {
	return creditcard.CoreCreditCard{
		Type:   data.Type,
		Name:   data.Name,
		Number: data.Number,
		Cvv:    data.Cvv,
		Month:  data.Month,
		Year:   data.Year,
		User: creditcard.CoreUser{
			ID: data.UserID,
		},
	}
}
