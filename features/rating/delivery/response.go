package delivery

import (
	"time"

	"github.com/GROUP-1-FASE-3/Backend/features/rating"
)

type RatingResponse struct {
	ID        uint      `json:"id"`
	Rating    uint      `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

func fromCore(dataCore rating.CoreRating) RatingResponse {
	return RatingResponse{
		ID:        dataCore.ID,
		Rating:    dataCore.Rating,
		Comment:   dataCore.Comment,
		CreatedAt: dataCore.CreatedAt,
	}
}

// data dari core ke response
func fromCoreList(dataCore []rating.CoreRating) []RatingResponse {
	var dataResponse []RatingResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
