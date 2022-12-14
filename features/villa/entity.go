package villa

type CoreVilla struct {
	ID             uint
	Villa_Name     string
	Price          uint
	Description    string
	Address        string
	Villa_Images1  string
	Villa_Images2  string
	Villa_Images3  string
	Detail_Guest   uint
	Detail_Bedroom uint
	Detail_Bed     uint
	Detail_Bath    uint
	Detail_Kitchen string
	Detail_Pool    string
	Detail_Wifi    string
	User           CoreUser
	Rating         CoreRating
}

type CoreUser struct {
	ID        uint
	User_Name string
}

type CoreRating struct {
	ID     uint
	Rating uint
}

type ServiceInterface interface {
	Create(input CoreVilla) (err error)
	GetAll() (data []CoreVilla, err error)
	GetById(id int) (data CoreVilla, err error)
	UpdateVilla(input CoreVilla, id int) (err error)
	DeleteVilla(id int) (err error)
}

type RepositoryInterface interface {
	Create(input CoreVilla) (row int, err error)
	GetAll() (data []CoreVilla, err error)
	GetById(id int) (data CoreVilla, err error)
	UpdateVilla(input CoreVilla, id int) (err error)
	DeleteVilla(id int) (row int, err error)
}
