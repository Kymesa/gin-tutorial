package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model        // Agrega campos ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `json:"title" validate:"required,min=3,max=100"` // Título del libro
	Author     string `json:"author" validate:"required,min=3,max=50"` // Autor del libro
}
