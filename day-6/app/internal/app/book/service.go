package book

import (
	"agmc/internal/dto"
	"agmc/internal/factory"
	"agmc/internal/model"
	"agmc/internal/repository"
	"context"
)

type service struct {
	BookRepository repository.Book
}

type Service interface {
	GetBooks(ctx context.Context) ([]dto.Book, int, error)
	GetBookByID(ctx context.Context, ID int64) (*dto.Book, int, error)
	CreateBook(ctx context.Context, b dto.CreateBookRequest) (*dto.Book, int, error)
	DeleteBookByID(ctx context.Context, ID int64) (int, error)
	UpdateBookByID(ctx context.Context, ID int64, data dto.UpdateBookRequest) (*dto.Book, int, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (s *service) GetBooks(ctx context.Context) ([]dto.Book, int, error) {
	books, code, err := s.BookRepository.GetBooks(ctx)
	if err != nil {
		return nil, code, err
	}

	res := make([]dto.Book, len(books))

	for i, book := range books {
		res[i] = dto.Book{
			ID:   book.ID,
			Name: book.Name,
		}
	}

	return res, code, nil
}
func (s *service) GetBookByID(ctx context.Context, ID int64) (*dto.Book, int, error) {
	book, code, err := s.BookRepository.GetBookByID(ctx, ID)
	if err != nil {
		return nil, code, err
	}

	return &dto.Book{
		ID:   book.ID,
		Name: book.Name,
	}, code, nil
}
func (s *service) CreateBook(ctx context.Context, b dto.CreateBookRequest) (*dto.Book, int, error) {
	book, code, err := s.BookRepository.CreateBook(ctx, model.Book{
		Name: b.Name,
	})
	if err != nil {
		return nil, code, err
	}

	return &dto.Book{
		ID:   book.ID,
		Name: book.Name,
	}, code, nil
}
func (s *service) DeleteBookByID(ctx context.Context, ID int64) (int, error) {
	return s.BookRepository.DeleteBookByID(ctx, ID)

}
func (s *service) UpdateBookByID(ctx context.Context, ID int64, data dto.UpdateBookRequest) (*dto.Book, int, error) {
	book, code, err := s.BookRepository.UpdateBookByID(ctx, ID, model.Book{
		Name: data.Name,
	})
	if err != nil {
		return nil, code, err
	}

	return &dto.Book{
		ID:   book.ID,
		Name: book.Name,
	}, code, nil
}
