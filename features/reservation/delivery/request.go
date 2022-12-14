package delivery

import (
	"time"

	"github.com/GROUP-1-FASE-3/Backend/features/reservation"
)

type ReservationRequest struct {
	Villa_ID     uint   `json:"villa_id" form:"villa_id"`
	Start_date   string `json:"start_date" form:"start_date"`
	End_date     string `json:"end_date" form:"end_date"`
	UserID       int    `json:"user_id" form:"user_id"`
	CreditCardID uint   `json:"credit_card_id" form:"credit_card_id"`
}

func toCore(data ReservationRequest) reservation.ReservationCore {
	return reservation.ReservationCore{
		VillaID:      data.Villa_ID,
		Start_Date:   timeParsing(data.Start_date),
		End_Date:     timeParsing(data.End_date),
		UserID:       uint(data.UserID),
		CreditCardID: data.CreditCardID,
	}
}

func timeParsing(date string) (dateParsed time.Time) {
	layoutFormat := "2006-01-02"

	dateParsed, _ = time.Parse(layoutFormat, date)

	return dateParsed
}
