package user

type UserCore struct {
	ID        uint
	User_Name string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
}

type RepositoryInterface interface {
	Create(input UserCore) (row int, err error)
}

type ServiceInterface interface {
	Create(input UserCore) (err error)
}
