package delivery

import (
	"net/http"
	"strconv"

	"github.com/GROUP-1-FASE-3/Backend/middlewares"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"

	"github.com/GROUP-1-FASE-3/Backend/features/rating"
	"github.com/labstack/echo/v4"
)

type RatingDelivery struct {
	ratingService rating.ServiceInterface
}

func New(service rating.ServiceInterface, e *echo.Echo) {
	handler := &RatingDelivery{
		ratingService: service,
	}

	e.POST("/ratings", handler.Create, middlewares.JWTMiddleware())
	e.GET("/ratings", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/ratings/:id", handler.GetById, middlewares.JWTMiddleware())
	e.PUT("/ratings/:id", handler.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/ratings/:id", handler.DeleteRating, middlewares.JWTMiddleware())
}

// Create
func (delivery *RatingDelivery) Create(c echo.Context) error {
	ratingInput := RatingRequest{}
	errBind := c.Bind(&ratingInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(ratingInput)
	err := delivery.ratingService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// Get All
func (delivery *RatingDelivery) GetAll(c echo.Context) error {
	results, err := delivery.ratingService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all ratings", dataResponse))
}

// Get by (id)
func (delivery *RatingDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	IdRating, err := delivery.ratingService.GetById(id)
	dataResponse := fromCore(IdRating)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get ratings", dataResponse))
}

// Update
func (delivery *RatingDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	ratingInput := RatingRequest{}
	errBind := c.Bind(&ratingInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(ratingInput)
	errUpt := delivery.ratingService.UpdateRating(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

// Delete
func (delivery *RatingDelivery) DeleteRating(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	errDel := delivery.ratingService.DeleteRating(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete rating"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}
