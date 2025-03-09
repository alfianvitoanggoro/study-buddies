package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Model interface {
	BeforeCreate(tx *gorm.DB) (err error)
	TableName() string
}

func CreateAllModel(db *gorm.DB) error {
	model := []interface{}{
		&User{},
		&Teacher{},
	}

	if err := db.AutoMigrate(model...); err != nil {
		logrus.Error(err)
	}

	return nil
}
