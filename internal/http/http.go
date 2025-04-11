package http

import (
	"fmt"
	"net/http"
	"os"
	"time"

	docs "github.com/AlfianVitoAnggoro/study-buddies/docs"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/app/user"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/factory"
	"github.com/AlfianVitoAnggoro/study-buddies/pkg/kafka"
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

	// kafka
	type RegistrationRequest struct {
		StudentID string `json:"student_id" validate:"required"`
		ClassID   string `json:"class_id" validate:"required"`
	}

	type ScheduleRegistrationRequest struct {
		ScheduleID string `json:"schedule_id" validate:"required"`
		ClassID    string `json:"class_id" validate:"required"`
		MaterialID string `json:"material_id" validate:"required"`
	}

	e.POST("/class-register", func(c echo.Context) error {
		var req RegistrationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
		}

		if req.StudentID == "" || req.ClassID == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "StudentID and ClassID are required"})
		}

		msg := kafka.ClassRegistrationMessage{
			StudentID: req.StudentID,
			ClassID:   req.ClassID,
			Timestamp: time.Now().Unix(),
		}

		if err := kafka.PublishClassRegistration(msg); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to publish message to Kafka"})
		}

		return c.JSON(http.StatusOK, echo.Map{"message": "Successfully registered class"})
	})

	e.POST("/schedule-register", func(c echo.Context) error {
		var req ScheduleRegistrationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
		}
		msg := kafka.ScheduleRegistrationMessage{
			ScheduleID: req.ScheduleID,
			ClassID:    req.ClassID,
			MaterialID: req.MaterialID,
			Timestamp:  time.Now().Unix(),
		}

		if err := kafka.PublishScheduleRegistration(msg); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to publish message to Kafka"})
		}

		return c.JSON(http.StatusOK, echo.Map{"message": "Successfully registered class"})
	})
}
