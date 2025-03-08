package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// ENV
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		fmt.Println(err)
	}
	PORT := env["PORT"]
	
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Hello, test!",
		})
	})
	e.Logger.Fatal(e.Start(":"+PORT))
}