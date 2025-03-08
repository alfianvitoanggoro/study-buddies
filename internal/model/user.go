package model

import (
	"os"
	"time"

	"github.com/AlfianVitoAnggoro/study-buddies/internal/abstraction"
	"github.com/AlfianVitoAnggoro/study-buddies/pkg/util/date"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	Id           uint   `json:"-"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required,email" gorm:"index:idx_user_email,unique"`
	PasswordHash string `json:"-"`
	Password     string `json:"password" validate:"required" gorm:"-"`
	IsActive     bool   `json:"is_active"`
}

type UserEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	UserEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type UserFilter struct {
	Code *string `query:"code" filter:"ILIKE"`
	Name *string `query:"name" filter:"ILIKE"`
}

type UserFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	UserFilter
}

func (UserEntityModel) TableName() string {
	return "users"
}

func (m *UserEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()

	m.hashPassword()
	m.Password = ""
	return
}

func (m *UserEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = date.DateTodayLocal()
	m.UpdatedBy = &m.Context.Auth.ID
	return
}

func (m *UserEntityModel) hashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *UserEntityModel) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         m.ID,
		"first_name": m.FirstName,
		"last":       m.LastName,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}