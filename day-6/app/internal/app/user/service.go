package user

import (
	"agmc/internal/dto"
	"agmc/internal/factory"
	"agmc/internal/model"
	"agmc/internal/repository"
	"context"
)

type service struct {
	BookRepository repository.Book
	UserRepository repository.User
}

type Service interface {
	CreateUser(ctx context.Context, data dto.CreateUserRequest) (*dto.User, int, error)
	GetUsers(ctx context.Context) ([]dto.User, int, error)
	GetUserByID(ctx context.Context, ID int64) (*dto.User, int, error)
	DeleteUserByID(ctx context.Context, ID int64) (int, error)
	UpdateUserByID(ctx context.Context, id int64, data dto.UpdateUserRequest) (int, error)
	// GetUserByEmail(ctx context.Context, email string) (*dto.User, int, error)
	UserLogin(ctx context.Context, email string, password string) (string, int, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
		BookRepository: f.BookRepository,
	}
}

func (s *service) CreateUser(ctx context.Context, data dto.CreateUserRequest) (*dto.User, int, error) {
	user, code, err := s.UserRepository.CreateUser(ctx, model.User{
		Name:     data.Name,
		Password: data.Password,
		Email:    data.Email,
	})

	if err != nil {
		return nil, code, err
	}

	return &dto.User{
		ID:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, code, nil
}

func (s *service) GetUsers(ctx context.Context) ([]dto.User, int, error) {
	users, code, err := s.UserRepository.GetUsers(ctx)

	if err != nil {
		return nil, code, err
	}

	res := make([]dto.User, len(users))

	for i, user := range users {
		res[i] = dto.User{
			ID:        user.ID,
			Name:      user.Name,
			Password:  user.Password,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		}
	}

	return res, code, nil
}

func (s *service) GetUserByID(ctx context.Context, ID int64) (*dto.User, int, error) {
	user, code, err := s.GetUserByID(ctx, ID)
	if err != nil {
		return nil, code, err
	}

	return &dto.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, code, nil
}

func (s *service) DeleteUserByID(ctx context.Context, ID int64) (int, error) {
	return s.UserRepository.DeleteUserByID(ctx, ID)
}

func (s *service) UpdateUserByID(ctx context.Context, id int64, data dto.UpdateUserRequest) (int, error) {
	code, err := s.UserRepository.UpdateUserByID(ctx, id, model.User{
		Name:     data.Name,
		Password: data.Password,
	})

	if err != nil {
		return code, err
	}

	return code, nil
}

func (s *service) UserLogin(ctx context.Context, email string, password string) (string, int, error) {
	return s.UserRepository.UserLogin(ctx, email, password)
}
