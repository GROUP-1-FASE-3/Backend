package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/user"

type UserResponse struct {
	ID           uint   `json:"id"`
	User_Name    string `json:"user_name"`
	Email        string `json:"email"`
	Gender       string `json:"gender"`
	Phone_Number string `json:"phone_number"`
	User_images  string `json:"user_images"`
}

func fromCore(dataCore user.UserCore) UserResponse {
	return UserResponse{
		ID:           dataCore.ID,
		User_Name:    dataCore.User_Name,
		Email:        dataCore.Email,
		Gender:       dataCore.Gender,
		Phone_Number: dataCore.Phone_Number,
		User_images:  dataCore.User_Images,
	}
}
