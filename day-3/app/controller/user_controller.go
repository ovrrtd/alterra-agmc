package controller

import (
	"agmc/lib/database"
	"agmc/model"
	"agmc/util"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type User struct {
	ID        int64          `json:"id" `
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func CreateUser(c echo.Context) error {
	u := new(CreateUserRequest)
	if err := c.Bind(u); err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	if err := c.Validate(u); err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	user, code, err := database.CreateUser(c.Request().Context(), model.User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
	})

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return util.ResponseWithJSON(c, "success create user", User{
		ID:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, code)
}

func GetUsers(c echo.Context) error {
	users, code, err := database.GetUsers(c.Request().Context())

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	res := make([]User, len(users))

	for i, user := range users {
		res[i] = User{
			ID:        user.ID,
			Name:      user.Name,
			Password:  user.Password,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		}
	}

	return util.ResponseWithJSON(c, "success get all users", res, code)
}

func GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	user, code, err := database.GetUserByID(c.Request().Context(), int64(id))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return util.ResponseWithJSON(c, "success get user by id", User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, code)
}

func DeleteUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	code, err := database.DeleteUserByID(c.Request().Context(), int64(id))

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return util.ResponseWithJSON(c, "success delete user by id", nil, code)
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UpdateUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	u := new(UpdateUserRequest)
	if err := c.Bind(u); err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	if err := c.Validate(u); err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	code, err := database.UpdateUserByID(c.Request().Context(), int64(id), model.User{
		Name:     u.Name,
		Password: u.Password,
	})

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return util.ResponseWithJSON(c, "success update user by id", nil, code)
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c echo.Context) error {
	r := new(LoginRequest)
	if err := c.Bind(&r); err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	if err := c.Validate(r); err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	token, code, err := database.UserLogin(c.Request().Context(), r.Email, r.Password)

	if err != nil {
		return util.ResponseWithJSON(c, err.Error(), nil, code)
	}
	return util.ResponseWithJSON(c, "success login", LoginResponse{
		Token: token,
	}, code)
}
