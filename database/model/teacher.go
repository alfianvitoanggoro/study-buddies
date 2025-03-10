package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Name       string    `json:"name"`
	Skills     string    `json:"skills"`
	gorm.Model           // Adds CreatedAt, UpdatedAt, DeletedAt
}

// BeforeCreate hook to generate UUID
func (u *Teacher) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New() // Generate UUID in Go instead of using PostgreSQL
	return
}

func (Teacher) TableName() string {
	return "teacher"
}
