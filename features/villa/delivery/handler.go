package delivery

import (
	"net/http"
	"strconv"

	"github.com/GROUP-1-FASE-3/Backend/features/villa"
	"github.com/GROUP-1-FASE-3/Backend/middlewares"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type VillaDelivery struct {
	villaService villa.ServiceInterface
}

func New(service villa.ServiceInterface, e *echo.Echo) {
	handler := &VillaDelivery{
		villaService: service,
	}
	e.POST("/villas", handler.Create, middlewares.JWTMiddleware())
	e.GET("/villas", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/villas/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/villas/:id", handler.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/villas/:id", handler.DeleteVilla, middlewares.JWTMiddleware())
	e.GET("/villas/user", handler.GetAllByID, middlewares.JWTMiddleware())
}

// Post New Villa
func (delivery *VillaDelivery) Create(c echo.Context) error {
	villaInput := VillaRequest{}
	errBind := c.Bind(&villaInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(villaInput)
	err := delivery.villaService.Create(dataCore, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// Get All Villa (Homepage)
func (delivery *VillaDelivery) GetAll(c echo.Context) error {
	results, err := delivery.villaService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataResponse))
}

// Get by ID
func (delivery *VillaDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	IdVilla, err := delivery.villaService.GetById(id)
	dataResponse := fromCore(IdVilla)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get users", dataResponse))
}

// Update
func (delivery *VillaDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	villaInput := VillaRequest{}
	errBind := c.Bind(&villaInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(villaInput)
	errUpt := delivery.villaService.UpdateVilla(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

// get all villa data by user_id
func (delivery *VillaDelivery) GetAllByID(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)

	results, err := delivery.villaService.GetAllByID(user_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", dataResponse))
}

// delete villa
func (delivery *VillaDelivery) DeleteVilla(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	errDel := delivery.villaService.DeleteVilla(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete villa"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
