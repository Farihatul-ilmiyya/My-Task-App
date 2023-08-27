package data

import (
	"mia/my_task_app/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string
}

// mapping core to model
func MapCoreUsertoUser(coreUser user.CoreUser) User {
	return User{
		Name:        coreUser.Name,
		Email:       coreUser.Email,
		Password:    coreUser.Password,
		PhoneNumber: coreUser.PhoneNumber,
	}
}

// Mapping User to CoreUser
func MapUserToCoreUser(dataModel User) user.CoreUser {
	return user.CoreUser{
		ID:          dataModel.ID,
		Name:        dataModel.Name,
		Email:       dataModel.Email,
		Password:    dataModel.Password,
		PhoneNumber: dataModel.PhoneNumber,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}
}

// Mapping dari []User ke []CoreUser
func ListMapUserToCoreUser(users []User) []user.CoreUser {
	var coreUsers []user.CoreUser
	for _, u := range users {
		coreUser := user.CoreUser{
			ID:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			Password:    u.Password,
			PhoneNumber: u.PhoneNumber,
			CreatedAt:   u.CreatedAt,
		}
		coreUsers = append(coreUsers, coreUser)
	}
	return coreUsers
}
