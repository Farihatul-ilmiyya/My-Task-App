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
func (s *UserService) CreateUser(inputUser user.CoreUser) (uint, error) {
	errValidate := s.validate.Struct(inputUser)
	if errValidate != nil {
		return 0, errors.New("validation error" + errValidate.Error())
	}
	userID, err := s.userRepo.Insert(inputUser)
	if err != nil {
		return 0, err
	}
	return userID, nil

}

// GetAllUser implements user.UserServiceInterface.
func (s *UserService) GetAllUser() ([]user.CoreUser, error) {
	result, err := s.userRepo.SelectAll()
	return result, err
}

// GetUserById implements user.UserServiceInterface.
func (s *UserService) GetUserById(userId uint) (user.CoreUser, error) {
	result, err := s.userRepo.Select(userId)
	return result, err
}

// DeleteUserById implements user.UserServiceInterface.
func (*UserService) DeleteUserById(userId uint) error {
	panic("unimplemented")
}

// Login implements user.UserServiceInterface.
func (*UserService) Login(email string, password string) (dataLogin user.CoreUser, token string, err error) {
	panic("unimplemented")
}

// UpdateUserById implements user.UserServiceInterface.
func (*UserService) UpdateUserById(userId uint, updateData user.CoreUser) error {
	panic("unimplemented")
}
