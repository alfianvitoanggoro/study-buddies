package seeder

import (
	"github.com/AlfianVitoAnggoro/study-buddies/database/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserSeeder seeds the users table
type TeacherSeeder struct{}

func (TeacherSeeder) Run(db *gorm.DB) error {
	teachers := []model.Teacher{
		{ID: uuid.New(), Name: "Alice", Skills: "Mathematics"},
		{ID: uuid.New(), Name: "Bob", Skills: "Physics"},
		{ID: uuid.New(), Name: "Charlie", Skills: "Chemistry"},
	}

	// Batch insert dengan satu query
	if err := db.Create(&teachers).Error; err != nil {
		return err
	}

	return nil
}
