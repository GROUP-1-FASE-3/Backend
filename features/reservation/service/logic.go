package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/GROUP-1-FASE-3/Backend/features/reservation"
	"github.com/go-playground/validator/v10"
)

type reservationService struct {
	reservationRepository reservation.RepositoryInterface
	validate              *validator.Validate
}

func New(repo reservation.RepositoryInterface) reservation.ServiceInterface {
	return &reservationService{
		reservationRepository: repo,
		validate:              validator.New(),
	}
}

func (s *reservationService) Check(input reservation.ReservationCore) (availability bool, err error) {
	if input.Start_Date.Unix() < time.Now().Unix() || input.End_Date.Unix() < time.Now().Unix() {
		return false, nil
	}

	if input.Start_Date.Unix() > input.End_Date.Unix() {
		return false, nil
	}

	listDate, err := s.reservationRepository.Check(input.VillaID)
	fmt.Println(listDate)
	if err != nil {
		return false, errors.New("error get data")
	}

	if listDate == nil {
		return true, nil
	}

	for _, date := range listDate {
		if (input.Start_Date.Unix() >= date.Start_Date.Unix() && input.Start_Date.Unix() <= date.End_Date.Unix()) || (input.End_Date.Unix() >= date.Start_Date.Unix() && input.End_Date.Unix() <= date.End_Date.Unix()) {
			return false, nil
		}
	}
	return true, nil
}

func (s *reservationService) Create(data reservation.ReservationCore) (err error) {
	if errValidate := s.validate.Struct(data); errValidate != nil {
		return errValidate
	}

	_, errCreate := s.reservationRepository.Create(data)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

func (s *reservationService) Get(id int) (data []reservation.ReservationCore, err error) {
	dataCore, errData := s.reservationRepository.Get(id)
	if errData != nil {
		return data, errData
	}

	// for _, v := range dataCore {
	// jmlMalam := v.End_Date.Sub(v.Start_Date)
	// 	v.Total_price = 4 * v.Price
	// }

	return dataCore, nil
}
