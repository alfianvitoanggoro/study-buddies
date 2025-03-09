package seeder

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Seeder interface ensures all seeders follow the same structure
type Seeder interface {
	Run(db *gorm.DB) error
}

// RunAllSeeders executes all seeders
func CreateAllSeeder(db *gorm.DB) error {
	seeders := []Seeder{
		UserSeeder{},
		TeacherSeeder{},
	}

	for _, s := range seeders {
		if err := s.Run(db); err != nil {
			logrus.Error("Seeding failed: ", err)
		}
	}

	return nil
}
