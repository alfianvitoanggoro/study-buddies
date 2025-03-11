package user

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetAllUsers)
	g.GET("/:id", h.GetUserByID)
	g.GET("/email/:email", h.GetUserByEmail)
	g.POST("", h.CreateUser)
	g.PUT("/:id", h.UpdateUser)
}
