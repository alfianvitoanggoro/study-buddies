package abstraction

import (
	"time"

	"github.com/AlfianVitoAnggoro/study-buddies/pkg/util/date"

	"gorm.io/gorm"
)

// Base entity model
type Entity struct {
	ID        string         `json:"id" gorm:"type:char(32);not null;primaryKey;index:unique;"`
	CreatedAt time.Time      `json:"createdAt"`
	CreatedBy *int           `json:"createdBy,omitempty"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty"`
	UpdatedBy *int           `json:"updatedBy,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Base filter model
type Filter struct {
	CreatedAt *time.Time `query:"createdAt"`
	CreatedBy *int       `query:"createdBy"`
	UpdatedAt *time.Time `query:"updatedAt"`
	UpdatedBy *int       `query:"updatedBy"`
}

// BeforeCreate: Set nilai default sebelum insert
func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	now := date.DateTodayLocal()
	m.CreatedAt = *now

	// Jika CreatedBy kosong, isi dengan default (misalnya Admin ID = 1)
	if m.CreatedBy == nil {
		defaultID := 1
		m.CreatedBy = &defaultID
	}

	return nil
}

// BeforeUpdate: Set updatedAt sebelum update
func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	now := date.DateTodayLocal()
	m.UpdatedAt = now

	// Jika UpdatedBy kosong, isi dengan default
	if m.UpdatedBy == nil {
		defaultID := 1
		m.UpdatedBy = &defaultID
	}

	return nil
}