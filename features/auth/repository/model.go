package repository

import "github.com/GROUP-1-FASE-3/Backend/features/auth"

type User struct {
	ID          uint
	User_Name   string
	Email       string `gorm:"unique"`
	Password    string
	User_Images string
}

func ToCore(userModel User) auth.Core {
	return auth.Core{
		ID:        userModel.ID,
		Email:     userModel.Email,
		User_name: userModel.User_Name,
		Password:  userModel.Password,
	}
}
