package repository

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/villa"
	"gorm.io/gorm"
)

type villaRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) villa.RepositoryInterface {
	return &villaRepository{
		db: db,
	}
}

func (repo *villaRepository) Create(input villa.CoreVilla) (row int, err error) {
	villaGorm := fromCore(input)
	tx := repo.db.Create(&villaGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements villa.Repository
func (repo *villaRepository) GetAll() (data []villa.CoreVilla, err error) {
	var villas []Villa

	tx := repo.db.Find(&villas)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(villas)
	return dataCore, nil
}
