package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/rating"

type RatingRequest struct {
	Rating  uint   `json:"rating" form:"rating"`
	Comment string `json:"comment" form:"comment"`
	VillaID uint   `json:"villa_id" form:"villa_id"`
	UserID  uint   `json:"user_id" form:"user_id"`
}

func toCore(data RatingRequest) rating.CoreRating {
	return rating.CoreRating{
		Rating:  data.Rating,
		Comment: data.Comment,
		User: rating.CoreUser{
			ID: data.UserID,
		},
		Villa: rating.CoreVilla{
			ID: data.VillaID,
		},
	}
}
