package repository

import (
	"github.com/GROUP-1-FASE-3/Backend/features/rating"
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Rating  uint
	Comment string
	VillaID uint
	UserID  uint
	User    User
	Villa   Villa
}

type User struct {
	gorm.Model
	User_Name    string
	Email        string
	Password     string
	Gender       string
	Phone_Number string
	User_Images  string
	Ratings      []Rating
}

type Villa struct {
	gorm.Model
	Villa_Name     string
	Price          string
	Description    string
	Address        string
	Villa_Images1  string
	Villa_Images2  string
	Villa_Images3  string
	Detail_Guest   uint
	Detail_Bedroom uint
	Detail_Bed     uint
	Detail_Bath    uint
	Detail_Kitchen string
	Detail_Pool    string
	Detail_Wifi    string
	UserID         uint
	Ratings        []Rating
}

// mengubah struct core ke struct model gorm
func fromCore(dataCore rating.CoreRating) Rating {
	ratingGorm := Rating{
		Rating:  dataCore.Rating,
		Comment: dataCore.Comment,
		VillaID: dataCore.Villa.ID,
		UserID:  dataCore.User.ID,
	}
	return ratingGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Rating) toCore() rating.CoreRating {
	return rating.CoreRating{
		ID:        dataModel.ID,
		Rating:    dataModel.Rating,
		Comment:   dataModel.Comment,
		CreatedAt: dataModel.CreatedAt,
		User: rating.CoreUser{
			ID:        dataModel.User.ID,
			User_Name: dataModel.User.User_Name,
		},
		Villa: rating.CoreVilla{
			ID:         dataModel.Villa.ID,
			Villa_Name: dataModel.Villa.Villa_Name,
		},
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Rating) []rating.CoreRating {
	var dataCore []rating.CoreRating
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
