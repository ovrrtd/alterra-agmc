package user

import (
	"agmc/internal/dto"
	"agmc/internal/factory"
	"agmc/pkg/util/response"
	"fmt"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) CreateUser(c echo.Context) error {
	fmt.Println("masuk")
	u := new(dto.CreateUserRequest)
	if err := c.Bind(u); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	if err := c.Validate(u); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	user, code, err := h.service.CreateUser(c.Request().Context(), *u)

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success create user", user, code)
}

func (h *handler) GetUsers(c echo.Context) error {
	users, code, err := h.service.GetUsers(c.Request().Context())

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success get all users", users, code)
}

func (h *handler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	user, code, err := h.service.GetUserByID(c.Request().Context(), int64(id))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success get user by id", user, code)
}

func (h *handler) DeleteUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	code, err := h.service.DeleteUserByID(c.Request().Context(), int64(id))

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success delete user by id", nil, code)
}

func (h *handler) UpdateUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	u := new(dto.UpdateUserRequest)
	if err := c.Bind(u); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	if err := c.Validate(u); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	code, err := h.service.UpdateUserByID(c.Request().Context(), int64(id), *u)

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}

	return response.ResponseWithJSON(c, "success update user by id", nil, code)
}

func (h *handler) Login(c echo.Context) error {
	r := new(dto.LoginRequest)
	if err := c.Bind(&r); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	if err := c.Validate(r); err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, http.StatusBadRequest)
	}

	token, code, err := h.service.UserLogin(c.Request().Context(), r.Email, r.Password)

	if err != nil {
		return response.ResponseWithJSON(c, err.Error(), nil, code)
	}
	return response.ResponseWithJSON(c, "success login", dto.LoginResponse{
		Token: token,
	}, code)
}
