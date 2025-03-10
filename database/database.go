package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	// configure database with PostgresSQL
	dbConfigurations := &DBPostgreSQL{
		DB: DB{
			Host: os.Getenv("DB_HOST"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASSWORD"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		},
		SslMode:  os.Getenv("DB_SSL_MODE"),
		TimeZone: os.Getenv("DB_TIME_ZONE"),
	}

	db, err := dbConfigurations.DBConnect()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (dbPgsl *DBPostgreSQL) DBConnect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbPgsl.Host, dbPgsl.User, dbPgsl.Pass, dbPgsl.Name, dbPgsl.Port, dbPgsl.SslMode, dbPgsl.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
