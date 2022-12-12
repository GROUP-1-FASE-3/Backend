package auth

type Core struct {
	ID        uint
	User_name string
	Email     string
	Password  string
	Token     string
}

type ServiceInterface interface {
	Login(email, password string) (loginData Core, err error)
}

type RepositoryInterface interface {
	FindUser(email, password string) (loginData Core, err error)
}
