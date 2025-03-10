package user

import (
	"github.com/AlfianVitoAnggoro/study-buddies/database/model"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/dto"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/factory"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type Service interface {
	GetAllUsers(c echo.Context, payload *dto.UserRequest) (*dto.UserResponse, error)
}

type service struct {
	Db             *gorm.DB
	RepositoryUser repository.User
}

func NewService(f *factory.Factory) *service {
	db := f.Db
	repositoryUser := f.UserRepository
	return &service{db, repositoryUser}
}

func (s *service) GetAllUsers() ([]dto.UserResponse, error) {
	var result []dto.UserResponse

	response, err := s.RepositoryUser.Find()
	if err != nil {
		return nil, err
	}

	for _, v := range *response {
		user := dto.UserResponse{
			Email: v.Email,
			Name:  v.Name,
		}
		result = append(result, user)
	}

	return result, nil
}

func (s *service) GetUserByEmail(email string) (*dto.UserByEmailResponse, error) {
	var result *dto.UserByEmailResponse

	response, err := s.RepositoryUser.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	result = &dto.UserByEmailResponse{
		Email:    response.Email,
		Name:     response.Name,
		Password: response.Password,
	}

	return result, nil
}

func (s *service) CreateUser(payload *dto.UserRequest) (*dto.UserResponse, error) {
	var result *dto.UserResponse

	id := uuid.UUID{}
	requestModel := &model.User{
		ID:       id,
		Email:    payload.Email,
		Name:     payload.Name,
		Password: payload.Password,
	}

	response, err := s.RepositoryUser.Create(requestModel)
	if err != nil {
		return nil, err
	}

	result = &dto.UserResponse{
		Email: response.Email,
		Name:  response.Name,
	}

	return result, nil
}
