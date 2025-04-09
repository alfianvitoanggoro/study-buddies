package http

import (
	"fmt"
	"net/http"
	"os"

	docs "github.com/AlfianVitoAnggoro/study-buddies/docs"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/app/user"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/factory"
	"github.com/AlfianVitoAnggoro/study-buddies/pkg/rabbitmq"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
		HOST    = os.Getenv("HOST")
		SCHEME  = os.Getenv("SCHEME")
	)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// doc
	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = HOST
	docs.SwaggerInfo.Schemes = []string{SCHEME}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	user.NewHandler(f).Route(e.Group("/user"))

	// rabbitmq
	e.POST("/create-class", func(c echo.Context) error {
		type Req struct {
			ClassName string `json:"class_name"`
			StudentID string `json:"student_id"`
		}

		var req Req
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		message := "Kelas baru: " + req.ClassName + " untuk murid ID " + req.StudentID
		rabbitmq.PublishMessage("class_notification", message)

		return c.JSON(http.StatusOK, echo.Map{"message": "Class created & notified!"})
	})
}
