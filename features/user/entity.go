package user

import "time"

type CoreUser struct {
	ID          uint
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type UserDataInterface interface {
	Login(email string, password string) (dataLogin CoreUser, string, err error)
	Insert(inputUser CoreUser) (uint, error)
	SelectAll() ([]CoreUser, error)
	Select(userId uint) (CoreUser, error)
	Update(userId uint, updateData CoreUser) error
	Delete(userId uint) error
}

type UserServiceInterface interface {
	Login(email string, password string) (dataLogin CoreUser, token string, err error)
	CreateUser(inputUser CoreUser) (uint, error)
	GetAllUser() ([]CoreUser, error)
	GetUserById(userId uint) (CoreUser, error)
	UpdateUserById(userId uint, updateData CoreUser) error
	DeleteUserById(userId uint) error
}
