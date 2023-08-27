package handler

import (
	"mia/my_task_app/features/user"
	"time"
)

type LoginResponse struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	// Role  string `json:"role"`
}
type UserResponse struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
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

// mapping from userCore to LoginResponse
func MapCoreUserToLogRes(Core user.CoreUser, jwtToken string) LoginResponse {
	return LoginResponse{
		Id:    Core.ID,
		Email: Core.Email,
		Token: jwtToken,
	}
}
