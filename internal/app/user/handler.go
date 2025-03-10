package user

import (
	"github.com/AlfianVitoAnggoro/study-buddies/internal/dto"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/factory"
	"github.com/AlfianVitoAnggoro/study-buddies/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

func NewHandler(f *factory.Factory) *handler {
	service := NewService(f)
	return &handler{service}
}

// GetAllUsers lists all existing users
//
//	@Summary		List users
//	@Description	get users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.UserResponse
//	@Failure		422	{object}	response.ErrorResponse
//	@Router			/user [get]
func (h *handler) GetAllUsers(c echo.Context) error {
	result, err := h.service.GetAllUsers()
	if err != nil {
		return c.JSON(422, response.ErrorResponse{
			Code:    422,
			Message: err.Error(),
		})
	}

	return c.JSON(200, result)
}

// GetUserByEmail get user by email
//
//	@Summary		Get user
//	@Description	get user by email
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			email	path		string	true	"User email"
//	@Success		200		{object}	dto.UserByEmailResponse
//	@Failure		402		{object}	response.ErrorResponse
//	@Failure		422		{object}	response.ErrorResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/user/{email} [get]
func (h *handler) GetUserByEmail(c echo.Context) error {

	payload := new(dto.UserByEmailRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(400, response.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(422, response.ErrorResponse{
			Code:    422,
			Message: err.Error(),
		})
	}

	result, err := h.service.GetUserByEmail(payload.Email)
	if err != nil {
		return c.JSON(422, response.ErrorResponse{
			Code:    422,
			Message: err.Error(),
		})
	}

	return c.JSON(200, result)
}

// CreateUser create user
//
//	@Summary		Create user
//	@Description	Create user to database
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserRequest	true	"request body"
//	@Success		200		{object}	dto.UserResponse
//	@Failure		402		{object}	response.ErrorResponse
//	@Failure		422		{object}	response.ErrorResponse
//	@Failure		500		{object}	response.ErrorResponse
//	@Router			/user [post]
func (h *handler) CreateUser(c echo.Context) error {

	payload := new(dto.UserRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(402, response.ErrorResponse{
			Code:    402,
			Message: err.Error(),
		})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(402, response.ErrorResponse{
			Code:    402,
			Message: err.Error(),
		})
	}

	result, err := h.service.CreateUser(payload)
	if err != nil {
		return c.JSON(422, response.ErrorResponse{
			Code:    422,
			Message: err.Error(),
		})
	}

	return c.JSON(200, result)
}
