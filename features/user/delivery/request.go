package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/user"

type UserRequest struct {
	User_Name string `json:"user_name" form:"full_name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

func toCore(data UserRequest) user.UserCore {
	return user.UserCore{
		User_Name: data.User_Name,
		Email:     data.Email,
		Password:  data.Password,
	}
}
