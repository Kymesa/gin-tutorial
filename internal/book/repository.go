package book

import (
	"gin-tutorial/internal/database"
)

type IRepository interface {
	Create(book *Book) error
	FindAll() ([]Book, error)
	FindByID(id uint) (*Book, error)
	Update(book *Book) error
	Delete(id uint) error
}

type bookRepository struct{}

func NewRepository() IRepository {
	return &bookRepository{}
}

// Create inserta un nuevo libro en la base de datos
func (r *bookRepository) Create(book *Book) error {
	return database.DB.Create(book).Error
}

// FindAll devuelve todos los libros
func (r *bookRepository) FindAll() ([]Book, error) {
	var books []Book
	err := database.DB.Find(&books).Error
	return books, err
}

// FindByID busca un libro por su ID
func (r *bookRepository) FindByID(id uint) (*Book, error) {
	var book Book
	err := database.DB.First(&book, id).Error
	return &book, err
}

// Update guarda los cambios de un libro existente
func (r *bookRepository) Update(book *Book) error {
	return database.DB.Save(book).Error
}

// Delete elimina un libro por ID
func (r *bookRepository) Delete(id uint) error {
	return database.DB.Delete(&Book{}, id).Error
}
