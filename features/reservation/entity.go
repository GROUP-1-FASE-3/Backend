package reservation

import "time"

type ReservationCore struct {
	ID           uint
	Start_Date   time.Time `validate:"required"`
	End_Date     time.Time `validate:"required"`
	Price        uint
	Total_price  uint
	VillaID      uint `validate:"required"`
	Villa        VillaCore
	UserID       uint
	CreditCardID uint
}

type DateReservation struct {
	Start_Date time.Time
	End_Date   time.Time
}

type VillaCore struct {
	Villa_Name string
	Price      uint
}

type RepositoryInterface interface {
	Check(id uint) (dataReservation []DateReservation, err error)
	Create(data ReservationCore) (row int, err error)
	Get(id int) (data []ReservationCore, err error)
}

type ServiceInterface interface {
	Check(data ReservationCore) (availability bool, err error)
	Create(data ReservationCore) (err error)
	Get(id int) (data []ReservationCore, err error)
}
