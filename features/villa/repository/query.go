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

	tx := repo.db.Preload("Ratings").Find(&villas)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(villas)
	return dataCore, nil
}

// Get by Id
func (repo *villaRepository) GetById(id int) (data villa.CoreVilla, err error) {
	var IdVilla Villa
	var IdVillaCore = villa.CoreVilla{}
	IdVilla.ID = uint(id)
	tx := repo.db.Preload("Ratings").First(&IdVilla, IdVilla.ID)
	if tx.Error != nil {
		return IdVillaCore, tx.Error
	}
	IdVillaCore = IdVilla.toCore()
	return IdVillaCore, nil
}

// Update
func (repo *villaRepository) UpdateVilla(datacore villa.CoreVilla, id int) (err error) {
	villaGorm := fromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(villaGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update villa failed")
	}
	return nil
}

// Delete
func (repo *villaRepository) DeleteVilla(id int) (row int, err error) {
	idVilla := Villa{}

	tx := repo.db.Delete(&idVilla, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete villa by id, failed")
	}
	return int(tx.RowsAffected), nil
}

func (repo *villaRepository) GetAllByID(id int) (data []villa.CoreVilla, err error) {
	var villas []Villa

	tx := repo.db.Preload("User").Where("user_id = ?", id).Find(&villas)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(villas)
	return dataCore, nil
}
