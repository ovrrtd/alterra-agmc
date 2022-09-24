package factory

import (
	"agmc/database"
	"agmc/internal/repository"
)

type Factory struct {
	BookRepository repository.Book
	UserRepository repository.User
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		BookRepository: repository.NewBook(),
		UserRepository: repository.NewUser(db),
	}
}
