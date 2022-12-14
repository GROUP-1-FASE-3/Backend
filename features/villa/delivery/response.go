package delivery

import "github.com/GROUP-1-FASE-3/Backend/features/villa"

type VillaResponse struct {
	ID               uint           `json:"id"`
	Villa_Name       string         `json:"vila_name"`
	Price            uint           `json:"price"`
	Description      string         `json:"description"`
	Address          string         `json:"address"`
	Villa_Images1    string         `json:"villa_images1"`
	Villa_Images2    string         `json:"villa_images2"`
	Villa_Images3    string         `json:"villa_images3"`
	Detail_Guest     uint           `json:"detail_guest"`
	Detail_Bedroom   uint           `json:"detail_bedroom"`
	Detail_Bed       uint           `json:"detail_bed"`
	Detail_Bath      uint           `json:"detail_bath"`
	Detail_Kitchen   string         `json:"detail_kitchen"`
	Detail_Pool      string         `json:"detail_pool"`
	Detail_Wifi      string         `json:"detail_wifi"`
	Detail_Workspace string         `json:"detail_workspace"`
	UserID           uint           `json:"user_id"`
	Ratings          RatingResponse `json:"ratings"`
}

type VillaResponseU struct {
	ID             uint         `json:"id"`
	Villa_Name     string       `json:"vila_name"`
	Price          uint         `json:"price"`
	Description    string       `json:"description"`
	Address        string       `json:"address"`
	Villa_Images1  string       `json:"villa_images1"`
	Villa_Images2  string       `json:"villa_images2"`
	Villa_Images3  string       `json:"villa_images3"`
	Detail_Guest   uint         `json:"detail_guest"`
	Detail_Bedroom uint         `json:"detail_bedroom"`
	Detail_Bed     uint         `json:"detail_bed"`
	Detail_Bath    uint         `json:"detail_bath"`
	Detail_Kitchen string       `json:"detail_kitchen"`
	Detail_Pool    string       `json:"detail_pool"`
	Detail_Wifi    string       `json:"detail_wifi"`
	Users          UserResponse `json:"users"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	User_Name string `json:"user_name"`
}

type RatingResponse struct {
	ID     uint `json:"id"`
	Rating uint `json:"rating"`
}

func fromCore(dataCore villa.CoreVilla) VillaResponse {
	return VillaResponse{
		ID:               dataCore.ID,
		Villa_Name:       dataCore.Villa_Name,
		Price:            dataCore.Price,
		Description:      dataCore.Description,
		Address:          dataCore.Address,
		Villa_Images1:    dataCore.Villa_Images1,
		Villa_Images2:    dataCore.Villa_Images2,
		Villa_Images3:    dataCore.Villa_Images3,
		Detail_Guest:     dataCore.Detail_Guest,
		Detail_Bedroom:   dataCore.Detail_Bedroom,
		Detail_Bed:       dataCore.Detail_Bed,
		Detail_Bath:      dataCore.Detail_Bath,
		Detail_Kitchen:   dataCore.Detail_Kitchen,
		Detail_Pool:      dataCore.Detail_Pool,
		Detail_Wifi:      dataCore.Detail_Wifi,
		Detail_Workspace: dataCore.Detail_Workspace,
		UserID:           dataCore.UserID,
		Ratings: RatingResponse{
			ID:     dataCore.Rating.ID,
			Rating: dataCore.Rating.Rating,
		},
	}
}

func fromCoreU(dataCore villa.CoreVilla) VillaResponseU {
	return VillaResponseU{
		ID:             dataCore.ID,
		Villa_Name:     dataCore.Villa_Name,
		Price:          dataCore.Price,
		Description:    dataCore.Description,
		Address:        dataCore.Address,
		Villa_Images1:  dataCore.Villa_Images1,
		Villa_Images2:  dataCore.Villa_Images2,
		Villa_Images3:  dataCore.Villa_Images3,
		Detail_Guest:   dataCore.Detail_Guest,
		Detail_Bedroom: dataCore.Detail_Bedroom,
		Detail_Bed:     dataCore.Detail_Bed,
		Detail_Bath:    dataCore.Detail_Bath,
		Detail_Kitchen: dataCore.Detail_Kitchen,
		Detail_Pool:    dataCore.Detail_Pool,
		Detail_Wifi:    dataCore.Detail_Wifi,
		Users: UserResponse{
			ID:        dataCore.User.ID,
			User_Name: dataCore.User.User_Name,
		},
	}
}

func fromCoreList(dataCore []villa.CoreVilla) []VillaResponse {
	var dataResponse []VillaResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreListU(dataCore []villa.CoreVilla) []VillaResponseU {
	var dataResponse []VillaResponseU
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreU(v))
	}
	return dataResponse
}
