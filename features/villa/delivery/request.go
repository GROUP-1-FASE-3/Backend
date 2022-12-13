package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/villa"

type VillaRequest struct {
	Villa_Name  string `json:"villa_name" form:"villa_name"`
	Price       uint   `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Address     string `json:"address" form:"address"`
	// Villa_Images1  string `json:"villa_images1" form:"villa_images1"`
	// Villa_Images2  string `json:"villa_images2" form:"villa_images2"`
	// Villa_Images3  string `json:"villa_images3" form:"villa_images3"`
	Detail_Guest   uint   `json:"detail_guest" form:"detail_guest"`
	Detail_Bedroom uint   `json:"detail_bedroom" form:"detail_bedroom"`
	Detail_Bed     uint   `json:"detail_bed" form:"detail_bed"`
	Detail_Bath    uint   `json:"detail_bath" form:"detail_bath"`
	Detail_Kitchen string `json:"detail_kitchen" form:"detail_kitchen"`
	Detail_Pool    string `json:"detail_pool" form:"detail_pool"`
	Detail_Wifi    string `json:"detail_wifi" form:"detail_wifi"`
	UserID         uint   `json:"user_id" form:"user_id"`
}

func toCore(data VillaRequest) villa.CoreVilla {
	return villa.CoreVilla{
		Villa_Name:  data.Villa_Name,
		Price:       data.Price,
		Description: data.Description,
		Address:     data.Address,
		// Villa_Images1:  data.Villa_Images1,
		// Villa_Images2:  data.Villa_Images2,
		// Villa_Images3:  data.Villa_Images3,
		Detail_Guest:   data.Detail_Guest,
		Detail_Bedroom: data.Detail_Bedroom,
		Detail_Bed:     data.Detail_Bed,
		Detail_Bath:    data.Detail_Bath,
		Detail_Kitchen: data.Detail_Kitchen,
		Detail_Pool:    data.Detail_Pool,
		Detail_Wifi:    data.Detail_Pool,
		User: villa.CoreUser{
			ID: data.UserID,
		},
	}
}
