package delivery

import (
	"net/http"

	"github.com/GROUP-1-FASE-3/Backend/features/user"
	"github.com/GROUP-1-FASE-3/Backend/middlewares"
	"github.com/GROUP-1-FASE-3/Backend/utils/helper"
	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}
	e.POST("/users", handler.Create, middlewares.JWTMiddleware())
}

func (d *UserDelivery) Create(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("operation failed"+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := d.userService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert user data"))
}
