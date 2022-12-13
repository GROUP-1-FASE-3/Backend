package villa

type CoreVilla struct {
	ID             uint
	Villa_Name     string
	Price          string
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
}

type CoreUser struct {
	ID        uint
	User_Name string
}

type ServiceInterface interface {
	Create(input CoreVilla) (err error)
}

type RepositoryInterface interface {
	Create(input CoreVilla) (row int, err error)
}
