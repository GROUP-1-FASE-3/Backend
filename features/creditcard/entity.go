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
	GetById(id int) (data CoreCreditCard, err error)
	UpdateCreditCard(input CoreCreditCard, id int) (err error)
	DeleteCreditCard(id int) (err error)
	GetByUserId(id int) (data []CoreCreditCard, err error)
}

type RepositoryInterface interface {
	Create(input CoreCreditCard) (row int, err error)
	GetById(id int) (data CoreCreditCard, err error)
	UpdateCreditCard(input CoreCreditCard, id int) (err error)
	DeleteCreditCard(id int) (row int, err error)
	GetByUserId(id int) (data []CoreCreditCard, err error)
}
