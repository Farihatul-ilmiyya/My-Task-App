package router

import (
	_userRepo "mia/my_task_app/features/user/data"
	_userHandler "mia/my_task_app/features/user/handler"
	_userService "mia/my_task_app/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepo := _userRepo.New(db)
	userService := _userService.New(userRepo)
	userHandlerAPI := _userHandler.New(userService)

	e.POST("/users", userHandlerAPI.CreateUser)
	e.GET("/users", userHandlerAPI.GetAllUser)
}