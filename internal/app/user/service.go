package user

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AlfianVitoAnggoro/study-buddies/database/model"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/dto"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/factory"
	"github.com/AlfianVitoAnggoro/study-buddies/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type Service interface {
	GetAllUsers(c echo.Context, payload *dto.UserRequest) (*dto.UserResponse, error)
}

type service struct {
	Db              *gorm.DB
	RedisRepository repository.RedisRepository
	RepositoryUser  repository.User
}

func NewService(f *factory.Factory) *service {
	db := f.Db
	repositoryRedis := f.RedisRepository
	repositoryUser := f.UserRepository
	return &service{db, repositoryRedis, repositoryUser}
}

func (s *service) GetAllUsers() ([]dto.UserResponse, error) {
	var result []dto.UserResponse
	var response []model.User // Mengubah menjadi slice langsung, bukan pointer ke slice

	// Get data dari Redis
	usersJson, err := s.RedisRepository.Get("getAllUsers")
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error when getting data from Redis: %s", err.Error()))
	}

	// Jika data ada di Redis, langsung unmarshall
	if usersJson != "" {
		if err := json.Unmarshal([]byte(usersJson), &response); err != nil {
			logrus.Error(fmt.Sprintf("Failed to unmarshal data from Redis: %s", err.Error()))
			response = nil // Pastikan response dikosongkan jika terjadi error
		}

		logrus.Info("Success when getting data from Redis")
	}

	// Jika tidak ada data di Redis atau terjadi error parsing, ambil dari database
	if len(response) == 0 {
		dbResponse, err := s.RepositoryUser.Find()
		if err != nil {
			return nil, fmt.Errorf("failed to fetch users from DB: %w", err)
		}

		// Gunakan nilai dari database
		response = *dbResponse

		// Simpan ke Redis dalam goroutine agar tidak memblokir permintaan utama
		go func(data []model.User) {
			jsonData, err := json.Marshal(data)
			if err != nil {
				logrus.Error(fmt.Sprintf("Failed to marshal users data: %s", err.Error()))
				return
			}

			if err := s.RedisRepository.Set("getAllUsers", string(jsonData), 10*time.Minute); err != nil {
				logrus.Error(fmt.Sprintf("Failed to store users data in Redis: %s", err.Error()))
			}

			logrus.Info("Success when storing users data in Redis")
		}(response)
	}

	// Konversi hasil ke DTO
	for _, v := range response {
		result = append(result, dto.UserResponse{
			Email: v.Email,
			Name:  v.Name,
		})
	}

	return result, nil
}

func (s *service) GetUserByID(id string) (*dto.UserByIDResponse, error) {
	var result *dto.UserByIDResponse
	var response *model.User
	key := fmt.Sprintf("GetUserByID:%s", id)

	// Get data dari Redis
	usersJson, err := s.RedisRepository.Get(key)
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error when getting data from Redis: %s", err.Error()))
	}

	// Jika data ada di Redis, langsung unmarshall
	if usersJson != "" {
		if err := json.Unmarshal([]byte(usersJson), &response); err != nil {
			logrus.Error(fmt.Sprintf("Failed to unmarshal data from Redis: %s", err.Error()))
			response = nil // Pastikan response dikosongkan jika terjadi error
		}

		logrus.Info("Success when getting data from Redis")
	}

	// Jika tidak ada data di Redis atau terjadi error parsing, ambil dari database
	if response == nil {
		dbResponse, err := s.RepositoryUser.FindByID(id)
		if err != nil {
			return nil, err
		}

		// Gunakan nilai dari database
		response = dbResponse

		// Simpan ke Redis dalam goroutine agar tidak memblokir permintaan utama
		go func(data model.User) {
			jsonData, err := json.Marshal(data)
			if err != nil {
				logrus.Error(fmt.Sprintf("Failed to marshal users data: %s", err.Error()))
				return
			}

			if err := s.RedisRepository.Set(key, string(jsonData), 10*time.Minute); err != nil {
				logrus.Error(fmt.Sprintf("Failed to store users data in Redis: %s", err.Error()))
			}

			logrus.Info("Success when storing users data in Redis")
		}(*response)
	}

	result = &dto.UserByIDResponse{
		Email:    response.Email,
		Name:     response.Name,
		Password: response.Password,
	}

	return result, nil
}

func (s *service) GetUserByEmail(email string) (*dto.UserByEmailResponse, error) {
	var result *dto.UserByEmailResponse
	var response *model.User
	key := fmt.Sprintf("GetUserByEmail:%s", email)

	// Get data dari Redis
	usersJson, err := s.RedisRepository.Get(key)
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error when getting data from Redis: %s", err.Error()))
	}

	// Jika data ada di Redis, langsung unmarshall
	if usersJson != "" {
		if err := json.Unmarshal([]byte(usersJson), &response); err != nil {
			logrus.Error(fmt.Sprintf("Failed to unmarshal data from Redis: %s", err.Error()))
			response = nil // Pastikan response dikosongkan jika terjadi error
		}

		logrus.Info("Success when getting data from Redis")
	}

	// Jika tidak ada data di Redis atau terjadi error parsing, ambil dari database
	if response == nil {
		dbResponse, err := s.RepositoryUser.FindByEmail(email)
		if err != nil {
			return nil, err
		}

		// Gunakan nilai dari database
		response = dbResponse

		// Simpan ke Redis dalam goroutine agar tidak memblokir permintaan utama
		go func(data model.User) {
			jsonData, err := json.Marshal(data)
			if err != nil {
				logrus.Error(fmt.Sprintf("Failed to marshal users data: %s", err.Error()))
				return
			}

			if err := s.RedisRepository.Set(key, string(jsonData), 10*time.Minute); err != nil {
				logrus.Error(fmt.Sprintf("Failed to store users data in Redis: %s", err.Error()))
			}

			logrus.Info("Success when storing users data in Redis")
		}(*response)
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

	go func() {
		// Menghapus data di Redis
		err := s.RedisRepository.Delete("getAllUsers")
		if err != nil {
			logrus.Warn(fmt.Sprintf("Error when deletting data after create from Redis: %s", err.Error()))
		}

		logrus.Info("Success when delete users data in Redis")
	}()

	result = &dto.UserResponse{
		Email: response.Email,
		Name:  response.Name,
	}

	return result, nil
}

func (s *service) UpdateUser(payload *dto.UserUpdateRequest) (*dto.UserUpdateResponse, error) {
	var result *dto.UserUpdateResponse

	requestModel := &model.User{
		Email:    payload.Email,
		Name:     payload.Name,
		Password: payload.Password,
	}

	response, err := s.RepositoryUser.Update(payload.ID, requestModel)
	if err != nil {
		return nil, err
	}

	go func() {
		// Menghapus all data di Redis
		err := s.RedisRepository.Delete("getAllUsers")
		if err != nil {
			logrus.Warn(fmt.Sprintf("Error when deletting data after create from Redis: %s", err.Error()))
		}

		logrus.Info("Success when delete users data in Redis")

		// Menghapus data di Redis by email
		err = s.RedisRepository.Delete("GetUserByID:" + payload.ID)
		if err != nil {
			logrus.Warn(fmt.Sprintf("Error when deletting data after create from Redis: %s", err.Error()))
		}

		logrus.Info("Success when delete user by ID data in Redis")

		if payload.Email == response.Email {
			err = s.RedisRepository.Delete("GetUserByEmail:" + payload.Email)
			if err != nil {
				logrus.Warn(fmt.Sprintf("Error when deletting data after create from Redis: %s", err.Error()))
			}
			logrus.Info("Success when delete user by email data in Redis")
		}

	}()

	result = &dto.UserUpdateResponse{
		Email:    response.Email,
		Name:     response.Name,
		Password: response.Password,
	}

	return result, nil
}
