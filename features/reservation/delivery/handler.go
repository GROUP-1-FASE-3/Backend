package delivery

import (
	"net/http"

	"github.com/GROUP-1-FASE-3/Backend/features/reservation"
	"github.com/GROUP-1-FASE-3/Backend/middlewares"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type ReservationDelivery struct {
	reservationService reservation.ServiceInterface
}

func New(service reservation.ServiceInterface, e *echo.Echo) {
	handler := &ReservationDelivery{
		reservationService: service,
	}

	e.POST("/reservations/check", handler.Check, middlewares.JWTMiddleware())
	e.POST("/reservations", handler.Create, middlewares.JWTMiddleware())
	e.GET("/reservations", handler.Get, middlewares.JWTMiddleware())
}

func (d *ReservationDelivery) Check(c echo.Context) error {
	dataReq := ReservationRequest{}
	c.Bind(&dataReq)

	dataCore := toCore(dataReq)

	availability, err := d.reservationService.Check(dataCore)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error check availability"))
	}

	if !availability {
		return c.JSON(http.StatusOK, helper.CheckResponse("not available"))
	}

	return c.JSON(http.StatusOK, helper.CheckResponse("available"))
}

func (d *ReservationDelivery) Create(c echo.Context) error {
	reserveInput := ReservationRequest{}
	errBind := c.Bind(&reserveInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("operation failed"+errBind.Error()))
	}

	reserveInput.UserID = middlewares.ExtractTokenUserId(c)

	dataCore := toCore(reserveInput)
	err := d.reservationService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert reservation data"))
}

func (d *ReservationDelivery) Get(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	result, err := d.reservationService.Get(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResp := fromCoreList(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get reservation data by user id", dataResp))
}
