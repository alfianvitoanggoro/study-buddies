package main

import (
	"context"
	"fmt"

	"github.com/AlfianVitoAnggoro/study-buddies/database"
	"github.com/AlfianVitoAnggoro/study-buddies/database/model"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/factory"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/http"
	"github.com/AlfianVitoAnggoro/study-buddies/pkg/cache"
	"github.com/AlfianVitoAnggoro/study-buddies/pkg/elasticsearch"
	"github.com/AlfianVitoAnggoro/study-buddies/pkg/util/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

//	@title			Study Buddies API
//	@version		0.0.1
//	@description	This is a documentation for Study Buddies API

// @host		localhost:8080
// @BasePath	/
func main() {
	// ENV
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		fmt.Println(err)
	}

	// Connect Database
	db, err := database.Init()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database %s", err.Error()))
	}

	logrus.Info(fmt.Sprintf("Successfully connected to database %s", db.Name()))

	// Create Model
	if err := model.CreateAllModel(db); err != nil {
		logrus.Error("Migration Model failed: ", err)
	}

	// Create Seeder Database
	// if err := seeder.CreateAllSeeder(db); err != nil {
	// 	logrus.Error("Seeding failed: ", err)
	// }

	ctx := context.Background()

	// Redis Connection
	rdb, err := cache.Init(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to redis database %s", err.Error()))
	}

	logrus.Info("Successfully connected to redis database")

	// Elastic Search Connection
	elasticsearch.Init()

	logrus.Info("Successfully connected to elastic search")

	e := echo.New()

	// Validate Request
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}

	f := factory.NewFactory(db, ctx, rdb)

	http.Init(e, f)

	e.Logger.Fatal(e.Start(":" + env["PORT"]))
}
