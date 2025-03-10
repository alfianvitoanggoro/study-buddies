package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Email      string    `json:"email" gorm:"index:idx_user_email,unique"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	gorm.Model           // Adds CreatedAt, UpdatedAt, DeletedAt
}

// BeforeCreate hook to generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New() // Generate UUID in Go instead of using PostgreSQL
	return
}

func (User) TableName() string {
	return "user"
}
