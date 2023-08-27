package handler

import (
	"mia/my_task_app/features/user"
	"time"
)

type UserResponse struct {
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	Password    string    `json:"password" form:"password"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
}

// mapping from userCore to UserResponse
func MapCoreUserToRes(Core user.CoreUser) UserResponse {
	return UserResponse{
		Name:        Core.Name,
		Email:       Core.Email,
		Password:    Core.Password,
		PhoneNumber: Core.PhoneNumber,
		CreatedAt:   Core.CreatedAt,
	}

}
