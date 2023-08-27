package handler

import (
	"mia/my_task_app/features/user"
	"mia/my_task_app/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	NewUser := new(UserRequest)

	// mendapatkan data yang dikirim oleh FE melalui request body
	err := c.Bind(&NewUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data", nil))
	}
	//mapping dari struct request to struct core
	input := ReqToCoreUser(*NewUser)
	_, err = h.userService.CreateUser(input)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))

		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success create user", nil))
}

func (h *UserHandler) GetAllUser(c echo.Context) error {
	result, err := h.userService.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var usersResponse []UserResponse
	for _, v := range result {
		usersResponse = append(usersResponse, MapCoreUserToRes(v))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusCreated, "success read data", usersResponse))
}
