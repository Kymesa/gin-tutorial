package book

import "gorm.io/gorm"

// Book representa la estructura del modelo que se mapeará a la tabla "books" en PostgreSQL
type Book struct {
	gorm.Model        // Agrega campos ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `json:"title"`  // Título del libro
	Author     string `json:"author"` // Autor del libro
}
