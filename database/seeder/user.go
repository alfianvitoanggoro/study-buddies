package seeder

import (
	"github.com/AlfianVitoAnggoro/study-buddies/database/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserSeeder seeds the users table
type UserSeeder struct{}

func (UserSeeder) Run(db *gorm.DB) error {
	users := []model.User{
		{ID: uuid.New(), Name: "Alice", Email: "alice@example.com", Password: "hashedpassword"},
		{ID: uuid.New(), Name: "Bob", Email: "bob@example.com", Password: "hashedpassword"},
		{ID: uuid.New(), Name: "Charlie", Email: "charlie@example.com", Password: "hashedpassword"},
	}

	for _, user := range users {
		db.Where("email = ?", user.Email).FirstOrCreate(&user)
	}

	return nil
}
