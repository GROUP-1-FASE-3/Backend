package repository

import (
	"errors"
	"time"

	"github.com/GROUP-1-FASE-3/Backend/features/reservation"
	"gorm.io/gorm"
)

type reserveRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.RepositoryInterface {
	return &reserveRepository{
		db: db,
	}
}

func (r *reserveRepository) Check(id uint) (dateReservation []reservation.DateReservation, err error) {
	var dates []reservation.DateReservation
	tx := r.db.Table("reservations").Select("start_date", "end_date").Where("villa_id = ? AND end_date > ?", id, time.Now()).Find(&dates)
	if tx.Error != nil {
		return nil, errors.New("failed read data")
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("failed read data")
	}

	return dates, nil
}

func (r *reserveRepository) Create(data reservation.ReservationCore) (row int, err error) {
	dataGorm := ReserveCoreToModel(data)

	tx := r.db.Create(&dataGorm)
	if tx.Error != nil {
		return -1, errors.New("failed create data")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("failed create data")
	}

	var id int
	ty := r.db.Raw("SELECT LAST_INSERT_ID()").Scan(&id)
	if ty.Error != nil {
		return -1, errors.New("failed create data")
	}

	tz := r.db.Exec("UPDATE reservations SET price = (SELECT price FROM villas WHERE id = ? ) WHERE id = ?", dataGorm.VillaID, id)
	if tz.Error != nil {
		return -1, errors.New("failed create data")
	}

	ta := r.db.Exec("UPDATE reservations SET total_price = (SELECT DATEDIFF(end_date, start_date) *price) WHERE id = ?", id)
	if ta.Error != nil {
		return -1, errors.New("failed create data")
	}
	return int(tx.RowsAffected), nil
}

func (r *reserveRepository) Get(id int) (data []reservation.ReservationCore, err error) {
	var dataReserve []Reservation

	tx := r.db.Where("user_id = ?", id).Preload("Villa").Find(&dataReserve)

	if tx.Error != nil {
		return data, nil
	}

	dataCore := toCoreList(dataReserve)
	return dataCore, nil
}
