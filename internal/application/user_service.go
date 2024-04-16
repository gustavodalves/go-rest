package application

import (
	"github.com/gustavodalves/go-api/internal/database"
	"github.com/gustavodalves/go-api/internal/entity"
)

type UserService struct {
	userDatabase *database.UserDB
}

func NewUserService(userDatabase *database.UserDB) *UserService {
	return &UserService{
		userDatabase: userDatabase,
	}
}

type RegisterNewUserDTO struct {
	Email    string
	Password string
}

func (s *UserService) Register(dto RegisterNewUserDTO) error {
	var user entity.User = entity.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	if err := s.userDatabase.Insert(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetAll() ([]*entity.User, error) {
	users, err := s.userDatabase.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetById(id uint64) (*entity.User, error) {
	user, err := s.userDatabase.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
