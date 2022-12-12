package delivery

import (
	"net/http"

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
}

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
