package user

import (
	"github.com/AlfianVitoAnggoro/study-buddies/internal/abstraction"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/dto"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/factory"

	"gorm.io/gorm"
)

type Service interface {
	GetAllUsers(ctx *abstraction.Context, payload *dto.UserRequest) (*dto.UserResponse, error)
}

type service struct {
	DB *gorm.DB
}

func NewService(f *factory.Factory) *service {
	db := f.Db
	return &service{db}
}

func (s *service) GetAllUsers(ctx *abstraction.Context, payload *dto.UserRequest) (*dto.UserResponse, error) {
	var result *dto.UserResponse

	response, err := s.db.FindAllUsers(ctx, payload)
	if err != nil {
		return nil, err
	}

	result = &dto.UserResponse{
		ID:    response.ID,
		Email: response.Email,
		Name:  response.Name,
	}

	return result, nil
}
