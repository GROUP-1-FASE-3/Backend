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
