package user

import (
	"github.com/AlfianVitoAnggoro/study-buddies/internal/abstraction"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/dto"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/factory"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

// CheckAccountBalance godoc
// @Summary CheckAccountBalance account
// @Description CheckAccountBalance account
// @Tags Account
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param request body dto.AccountCheckBalanceGetRequest true "request body"
// @Success 200 {object} dto.AccountCheckBalanceGetResponseDoc
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /account/checkBalance [get]
func (h *handler) GetAllUsers(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.UserRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err.Error())
	}

	result, err := h.service.GetAllUsers(cc, payload)
	if err != nil {
		return c.JSON(422, err.Error())
	}

	// return res.SuccessResponse(result).Send(c)
	return c.JSON(200, result)
}
