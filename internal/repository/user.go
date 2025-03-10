package repository

import (
	"github.com/AlfianVitoAnggoro/study-buddies/internal/abstraction"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type User interface {
	FindAllUsers(c echo.Context) error
}

type user struct {
	abstraction.Repository
}

func NewUser(db *gorm.DB) *user {
	return &user{
		abstraction.Repository{
			Db: db,
		},
	}
}

type UserRepository struct {
}

func (ur *UserRepository) FindAllUsers(c echo.Context) error {
	return nil
}
