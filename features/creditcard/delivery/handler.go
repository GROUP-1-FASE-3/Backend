package delivery

import (
	"net/http"
	"strconv"

	"github.com/GROUP-1-FASE-3/Backend/features/creditcard"
	"github.com/GROUP-1-FASE-3/Backend/middlewares"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type CreditCardDelivery struct {
	creditcardService creditcard.ServiceInterface
}

func New(service creditcard.ServiceInterface, e *echo.Echo) {
	handler := &CreditCardDelivery{
		creditcardService: service,
	}

	e.POST("/creditcards", handler.Create, middlewares.JWTMiddleware())
	e.GET("/creditcards/:id", handler.GetById, middlewares.JWTMiddleware())
	e.GET("/creditcards/user", handler.GetByUId, middlewares.JWTMiddleware())
	e.PUT("/creditcards/:id", handler.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/creditcards/:id", handler.DeleteCreditCard, middlewares.JWTMiddleware())
}

// Post
func (delivery *CreditCardDelivery) Create(c echo.Context) error {
	creditcardInput := CreditCardRequest{}
	errBind := c.Bind(&creditcardInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(creditcardInput)
	err := delivery.creditcardService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

// Get by (ID)
func (delivery *CreditCardDelivery) GetById(c echo.Context) error {
	id, errBind := strconv.Atoi(c.Param("id"))
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	IdCreditCard, err := delivery.creditcardService.GetById(id)
	dataResponse := fromCore(IdCreditCard)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get users", dataResponse))
}

// Update by (ID)
func (delivery *CreditCardDelivery) UpdateData(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	creditcardInput := CreditCardRequest{}
	errBind := c.Bind(&creditcardInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := toCore(creditcardInput)
	errUpt := delivery.creditcardService.UpdateCreditCard(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

// Delete
func (delivery *CreditCardDelivery) DeleteCreditCard(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error conv data "+errConv.Error()))
	}

	errDel := delivery.creditcardService.DeleteCreditCard(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete creditcard"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}

func (delivery *CreditCardDelivery) GetByUId(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	IdCreditCard, err := delivery.creditcardService.GetByUserId(id)
	dataResponse := fromCoreList(IdCreditCard)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get users", dataResponse))
}
