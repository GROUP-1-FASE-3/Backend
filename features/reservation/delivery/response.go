package delivery

import (
	"time"

	"github.com/GROUP-1-FASE-3/Backend/features/reservation"
)

type ReserveResponse struct {
	User_id        uint      `json:"user_id"`
	Reservation_id uint      `json:"reservation_id"`
	Villa_name     string    `json:"villa_name"`
	Start_date     time.Time `json:"start_date"`
	End_date       time.Time `json:"end_date"`
	Price          uint      `json:"price"`
	Total_price    uint      `json:"total_price"`
}

func fromCore(dataCore reservation.ReservationCore) ReserveResponse {
	return ReserveResponse{
		User_id:        dataCore.UserID,
		Reservation_id: dataCore.ID,
		Villa_name:     dataCore.Villa.Villa_Name,
		Start_date:     dataCore.Start_Date,
		End_date:       dataCore.End_Date,
		Price:          dataCore.Villa.Price,
		Total_price:    dataCore.Total_price,
	}
}

func fromCoreList(dataCore []reservation.ReservationCore) []ReserveResponse {
	var dataResponse []ReserveResponse

	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
