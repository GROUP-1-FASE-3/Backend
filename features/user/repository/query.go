package repository

import (
	"errors"

	"github.com/GROUP-1-FASE-3/Backend/features/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(input user.UserCore) (row int, err error) {
	userGorm := UserCoreToModel(input)
	tx := r.db.Create(&userGorm)
	if tx.Error != nil {
		return -1, errors.New("failed create data")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("failed create data")
	}
	return int(tx.RowsAffected), nil
}

func (r *userRepository) Delete(data user.UserCore, id int) (row int, err error) {
	userGorm := UserCoreToModel(data)

	tx := r.db.Delete(&userGorm, id)
	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("data not found")
	}

	return int(tx.RowsAffected), nil
}

func (r *userRepository) GetByID(id int) (data user.UserCore, err error) {
	var user User

	tx := r.db.First(&user, id)

	if tx.Error != nil {
		return data, errors.New("data not found")
	}

	userCore := user.toCore()

	return userCore, nil
}

func (r *userRepository) Update(id int, data user.UserCore) (row int, err error) {
	var user User
	gormUserCore := UserCoreToModel(data)

	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return -1, errors.New("data not found")
	}

	tz := r.db.Model(&user).Updates(gormUserCore)
	if tz.Error != nil {
		return -1, errors.New("data not found")
	}

	if tz.RowsAffected == 0 {
		return 0, err
	}

	return int(tz.RowsAffected), nil
}
