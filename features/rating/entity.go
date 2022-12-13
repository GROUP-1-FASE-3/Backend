package rating

import "time"

type CoreRating struct {
	ID        uint
	Rating    uint
	Comment   string
	VillaID   uint
	UserID    uint
	CreatedAt time.Time
	User      CoreUser
	Villa     CoreVilla
}

type CoreUser struct {
	ID        uint
	User_Name string
}

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
	UserID         uint
}

type ServiceInterface interface {
	GetAll() (data []CoreRating, err error)
	Create(input CoreRating) (err error)
	GetById(id int) (data CoreRating, err error)
	UpdateRating(input CoreRating, id int) (err error)
	DeleteRating(id int) (err error)
}

type RepositoryInterface interface {
	GetAll() (data []CoreRating, err error)
	Create(input CoreRating) (row int, err error)
	GetById(id int) (data CoreRating, err error)
	UpdateRating(input CoreRating, id int) (err error)
	DeleteRating(id int) (row int, err error)
}
