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

// Post
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

// Get by (ID)
func (repo *creditcardRepository) GetById(id int) (data creditcard.CoreCreditCard, err error) {
	var IdCreditCard CreditCard
	var IdCreditCardCore = creditcard.CoreCreditCard{}
	IdCreditCard.ID = uint(id)
	tx := repo.db.Preload("Users").First(&IdCreditCard, IdCreditCard.ID)
	if tx.Error != nil {
		return IdCreditCardCore, tx.Error
	}
	IdCreditCardCore = IdCreditCard.toCore()
	return IdCreditCardCore, nil
}

// Update
func (repo *creditcardRepository) UpdateCreditCard(datacore creditcard.CoreCreditCard, id int) (err error) {
	creditcardGorm := fromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(creditcardGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update creditcard failed")
	}
	return nil
}

// Delete
func (repo *creditcardRepository) DeleteCreditCard(id int) (row int, err error) {
	IdCreditCard := CreditCard{}

	tx := repo.db.Delete(&IdCreditCard, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete creditcard by id failed")
	}
	return int(tx.RowsAffected), nil
}
