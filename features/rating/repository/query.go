package repository

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/rating"
	"gorm.io/gorm"
)

type ratingRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) rating.RepositoryInterface {
	return &ratingRepository{
		db: db,
	}
}

// Create implements rating.Repository
func (repo *ratingRepository) Create(input rating.CoreRating) (row int, err error) {
	ratingGorm := fromCore(input)
	tx := repo.db.Create(&ratingGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *ratingRepository) GetAll() (data []rating.CoreRating, err error) {
	var ratings []Rating

	tx := repo.db.Find(&ratings)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(ratings)
	return dataCore, nil
}

// GetById implements rating.RepositoryInterface
func (repo *ratingRepository) GetById(id int) (data rating.CoreRating, err error) {
	var IdRating Rating
	var IdRatingCore = rating.CoreRating{}
	IdRating.ID = uint(id)
	tx := repo.db.First(&IdRating, IdRating.ID)
	if tx.Error != nil {
		return IdRatingCore, tx.Error
	}
	IdRatingCore = IdRating.toCore()
	return IdRatingCore, nil
}

// Update
func (repo *ratingRepository) UpdateRating(datacore rating.CoreRating, id int) (err error) {
	ratingGorm := fromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(ratingGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update rating failed")
	}
	return nil
}

// Delete
func (repo *ratingRepository) DeleteRating(id int) (row int, err error) {
	idRating := Rating{}

	tx := repo.db.Delete(&idRating, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete rating by id failed")
	}
	return int(tx.RowsAffected), nil
}
