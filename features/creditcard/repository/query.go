package repository

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/creditcard"
	"gorm.io/gorm"
)

type creditcardRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) creditcard.RepositoryInterface {
	return &creditcardRepository{
		db: db,
	}
}

func (repo *creditcardRepository) Create(input creditcard.CoreCreditCard) (row int, err error) {
	creditcardGorm := fromCore(input)
	tx := repo.db.Create(&creditcardGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}
