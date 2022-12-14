package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/user"

type UserRequest struct {
	User_Name    string `json:"user_name" form:"user_name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Gender       string `json:"gender" form:"gender"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
	User_Images  string `json:"user_images" form:"user_images"`
}

func toCore(data UserRequest) user.UserCore {
	return user.UserCore{
		User_Name:    data.User_Name,
		Email:        data.Email,
		Password:     data.Password,
		Gender:       data.Gender,
		Phone_Number: data.Phone_Number,
	}
}
