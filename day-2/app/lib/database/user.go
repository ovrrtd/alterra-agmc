package database

import (
	"context"
	"day-two/config"
	"day-two/model"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, user model.User) (*model.User, int, error) {
	if err := config.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &user, http.StatusCreated, nil
}

func GetUsers(ctx context.Context) ([]model.User, int, error) {
	var users []model.User
	if err := config.DB.WithContext(ctx).Find(&users).Where(`deleted_at IS NULL`).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return users, http.StatusOK, nil
}

func GetUserByID(ctx context.Context, ID int64) (*model.User, int, error) {
	var user model.User

	if err := config.DB.WithContext(ctx).First(&user, `id = ?`, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		return nil, http.StatusInternalServerError, err
	}

	return &user, http.StatusOK, nil
}

func DeleteUserByID(ctx context.Context, ID int64) (int, error) {
	result := config.DB.WithContext(ctx).Where(`id = ?`, ID).Delete(&model.User{})
	if err := result.Error; err != nil {
		return http.StatusInternalServerError, err
	}

	if result.RowsAffected == 0 {
		return http.StatusNotFound, errors.New("user not found")
	}

	return http.StatusOK, nil
}

func UpdateUserByID(ctx context.Context, id int64, data model.User) (int, error) {
	result := config.DB.WithContext(ctx).Model(&model.User{}).Where(`deleted_at IS NULL AND id = ?`, id).Updates(data)
	if err := result.Error; err != nil {
		return http.StatusInternalServerError, err
	}

	if result.RowsAffected == 0 {
		return http.StatusNotFound, errors.New("user not found")
	}

	return http.StatusOK, nil
}
