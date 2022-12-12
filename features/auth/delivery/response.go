package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/auth"

type LoginResponse struct {
	ID          uint   `json:"id"`
	User_name   string `json:"user_name"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func fromCore(dataCore auth.Core) LoginResponse {
	return LoginResponse{
		ID:          dataCore.ID,
		User_name:   dataCore.User_name,
		Email:       dataCore.Email,
		TokenString: dataCore.Token,
	}
}
