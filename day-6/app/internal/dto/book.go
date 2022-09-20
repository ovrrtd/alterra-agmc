package dto

type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type CreateBookRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateBookRequest struct {
	Name string `json:"name" validate:"required"`
}
