package services

import (
	"fmt"

	"github.com/gene-qxsi/CRM-M/internal/models"
	"github.com/gene-qxsi/CRM-M/internal/storage"
)

type UserService struct {
	Storage *storage.Storage
}

func New(storage *storage.Storage) *UserService {
	return &UserService{Storage: storage}
}

// TODO написать обработку возможных ошибок
func (s *UserService) CreateUser(user models.User) (int, error) {
	if user.Name == "" {
		return 0, fmt.Errorf("имя пользователя не должно быть пустым")
	}
	if len(user.Name) > 24 {
		return 0, fmt.Errorf("имя пользователя должно быть не больше 24 символов")
	}
	return s.Storage.CreateUser(user)
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id пользователя должно быть больше нуля")
	}
	return s.Storage.GetUser(id)
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.Storage.GetUsers()
}

func (s *UserService) DeleteUser(id int) error {
	return s.Storage.DeleteUser(id)
}

func (s *UserService) GetUserByNameAndPassword(name, password string) (*models.User, error) {
	return s.Storage.GetUserByNameAndPassword(name, password)
}
