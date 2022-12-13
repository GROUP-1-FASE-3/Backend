package delivery

import (
	"net/http"

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
}

// Post New Villa
func (delivery *VillaDelivery) Create(c echo.Context) error {
	villaInput := VillaRequest{}
	errBind := c.Bind(&villaInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(villaInput)
	err := delivery.villaService.Create(dataCore)
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
