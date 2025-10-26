package store

import (
	"fmt"
	"gin-tutorial/internal/model"

	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetAll() ([]model.Libro, error)
	GetById(id int) (*model.Libro, error)
	Create(book model.Libro) (model.Libro, error)
	Update(id int, libro model.Libro) (model.Libro, error)
	Delete(id int) error
}

type store struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Store {

	return &store{db: db}
}

func (s *store) GetAll() ([]model.Libro, error) {

	query := ` SELECT id, title, author FROM books `

	var books []model.Libro
	err := s.db.Select(&books, query)
	return books, err
}

func (s *store) GetById(id int) (*model.Libro, error) {

	query := ` SELECT id, title, author FROM books WHERE id = ? `

	book := model.Libro{}

	err := s.db.Get(&book, query, id)
	return &book, err
}

func (s *store) Create(book model.Libro) (model.Libro, error) {

	query := ` INSERT INTO books (title, author) VALUES (:title, :author) `

	res, err := s.db.NamedExec(
		query,
		book,
	)
	if err != nil {
		return model.Libro{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Libro{}, err
	}

	book.ID = int(id)

	return book, nil
}

func (s *store) Update(id int, book model.Libro) (model.Libro, error) {

	book.ID = id

	query := `
		UPDATE books
		SET title = :title, author = :author
		WHERE id = :id
	`

	result, err := s.db.NamedExec(query, book)

	if err != nil {
		return model.Libro{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Libro{}, err
	}

	if rowsAffected == 0 {
		return model.Libro{}, fmt.Errorf("no se encontró ningún libro con id %d", book.ID)
	}

	return book, nil
}

func (s *store) Delete(id int) error {

	query := ` DELETE FROM books WHERE id = :id `

	result, err := s.db.NamedExec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return err
	}

	return nil
}
