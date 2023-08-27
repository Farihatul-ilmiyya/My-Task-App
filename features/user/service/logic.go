package service

import (
	"errors"
	"mia/my_task_app/features/user"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userRepo user.UserDataInterface
	validate *validator.Validate
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &UserService{
		userRepo: repo,
		validate: validator.New(),
	}
}

// CreateUser implements user.UserServiceInterface.
func (s *UserService) Create(inputUser user.CoreUser) (uint, error) {
	errValidate := s.validate.Struct(inputUser)
	if errValidate != nil {
		return 0, errors.New("validation error, " + errValidate.Error())
	}
	userID, err := s.userRepo.Insert(inputUser)
	if err != nil {
		return 0, err
	}
	return userID, nil

}

// GetAllUser implements user.UserServiceInterface.
func (s *UserService) GetAll() ([]user.CoreUser, error) {
	result, err := s.userRepo.SelectAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserById implements user.UserServiceInterface.
func (s *UserService) GetById(userId uint) (user.CoreUser, error) {
	result, err := s.userRepo.Select(userId)
	if err != nil {
		return user.CoreUser{}, err
	}
	return result, nil
}

// UpdateUserById implements user.UserServiceInterface.
func (s *UserService) UpdateById(userId uint, userData user.CoreUser) error {
	err := s.userRepo.Update(userId, userData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserById implements user.UserServiceInterface.
func (s *UserService) DeleteById(userId uint) error {
	err := s.userRepo.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}

// Login implements user.UserServiceInterface.
func (*UserService) Login(email string, password string) (dataLogin user.CoreUser, token string, err error) {
	panic("unimplemented")
}
