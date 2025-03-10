package main

import (
	"fmt"

	"github.com/AlfianVitoAnggoro/study-buddies/database"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/http"
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

	// Connect Database
	database.Init()

	e := echo.New()

	http.Init(e, nil, nil)

	e.Logger.Fatal(e.Start(":" + env["PORT"]))
}
