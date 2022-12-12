package creditcard

type CoreCreditCard struct {
	ID     uint
	Type   string
	Name   string
	Number string
	Cvv    uint
	Month  uint
	Year   uint
	User   CoreUser
}

type CoreUser struct {
	ID        uint
	User_Name string
}

type ServiceInterface interface {
	Create(input CoreCreditCard) (err error)
}

type RepositoryInterface interface {
	Create(input CoreCreditCard) (row int, err error)
}
