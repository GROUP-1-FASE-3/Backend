package repository

import (
	"github.com/GROUP-1-FASE-3/Backend/features/auth"
	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) FindUser(email, password string) (loginData auth.Core, err error) {
	userModel := User{}
	tx := repo.db.Where("email= ? and deleted_at is null", email).First(&userModel)

	err = tx.Error

	if err != nil {
		return loginData, err
	}

	loginData = ToCore(userModel)
	return loginData, nil
}
