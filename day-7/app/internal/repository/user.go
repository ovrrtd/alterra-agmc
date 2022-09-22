package repository

import (
	"agmc/internal/middleware"
	"agmc/internal/model"
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, int, error)
	GetUsers(ctx context.Context) ([]model.User, int, error)
	GetUserByID(ctx context.Context, ID int64) (*model.User, int, error)
	DeleteUserByID(ctx context.Context, ID int64) (int, error)
	UpdateUserByID(ctx context.Context, id int64, data model.User) (int, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, int, error)
	UserLogin(ctx context.Context, email string, password string) (string, int, error)
}

type user struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (u *user) CreateUser(ctx context.Context, user model.User) (*model.User, int, error) {
	if err := u.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &user, http.StatusCreated, nil
}

func (u *user) GetUsers(ctx context.Context) ([]model.User, int, error) {
	var users []model.User
	if err := u.DB.WithContext(ctx).Find(&users).Where(`deleted_at IS NULL`).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return users, http.StatusOK, nil
}

func (u *user) GetUserByID(ctx context.Context, ID int64) (*model.User, int, error) {
	var user model.User

	if err := u.DB.WithContext(ctx).First(&user, `id = ?`, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}

func (u *user) DeleteUserByID(ctx context.Context, ID int64) (int, error) {
	result := u.DB.WithContext(ctx).Where(`id = ?`, ID).Delete(&model.User{})
	if err := result.Error; err != nil {
		return http.StatusInternalServerError, err
	}

	if result.RowsAffected == 0 {
		return http.StatusNotFound, errors.New("user not found")
	}

	return http.StatusOK, nil
}

func (u *user) UpdateUserByID(ctx context.Context, id int64, data model.User) (int, error) {
	result := u.DB.WithContext(ctx).Model(&model.User{}).Where(`deleted_at IS NULL AND id = ?`, id).Updates(data)
	if err := result.Error; err != nil {
		return http.StatusInternalServerError, err
	}

	if result.RowsAffected == 0 {
		return http.StatusNotFound, errors.New("user not found")
	}

	return http.StatusOK, nil
}

func (u *user) GetUserByEmail(ctx context.Context, email string) (*model.User, int, error) {
	var user model.User

	if err := u.DB.WithContext(ctx).First(&user, `email = ?`, email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}

func (u *user) UserLogin(ctx context.Context, email string, password string) (string, int, error) {
	user, code, err := u.GetUserByEmail(ctx, email)

	if err != nil {
		return "", code, err
	}

	if user.Password != password {
		return "", http.StatusBadRequest, errors.New("wrong password")
	}

	token, err := middleware.CreateToken(user.ID)
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("internal server error")
	}

	return token, http.StatusOK, nil
}
