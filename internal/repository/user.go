package repository

import (
	"github.com/AlfianVitoAnggoro/study-buddies/database/model"
	"gorm.io/gorm"
)

type User interface {
	Find() (*[]model.User, error)
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(payload *model.User) (*model.User, error)
	Update(id string, payload *model.User) (*model.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (u *user) Find() (*[]model.User, error) {
	var datas []model.User
	query := u.db.Model(&model.User{})

	err := query.Find(&datas).Error

	if err != nil {
		return &datas, err
	}

	return &datas, nil
}

func (u *user) FindByID(id string) (*model.User, error) {
	var data model.User
	query := u.db.Model(&model.User{})

	err := query.Where("id = ?", id).First(&data).Error

	if err != nil {
		return &data, err
	}

	return &data, nil
}

func (u *user) FindByEmail(email string) (*model.User, error) {
	var data model.User
	query := u.db.Model(&model.User{})

	err := query.Where("email = ?", email).First(&data).Error

	if err != nil {
		return &data, err
	}

	return &data, nil
}

func (u *user) Create(payload *model.User) (*model.User, error) {
	var data model.User

	query := u.db.Model(&model.User{})

	err := query.Create(&payload).Error

	if err != nil {
		return &data, err
	}

	return payload, nil
}

func (u *user) Update(id string, payload *model.User) (*model.User, error) {
	var data model.User

	query := u.db.Model(&model.User{})

	err := query.Where("id = ?", id).Updates(&payload).Error

	if err != nil {
		return &data, err
	}

	return payload, nil
}
