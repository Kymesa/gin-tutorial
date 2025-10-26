package services

import (
	"errors"
	"gin-tutorial/internal/model"
	"gin-tutorial/internal/store"
)

type Logger interface {
	Log(msg string, error string)
}

type Services struct {
	store  store.Store
	logger Logger
}

func New(s store.Store) *Services {
	return &Services{store: s, logger: nil}
}

func (s *Services) ObtenTodosLosLibros() ([]model.Libro, error) {
	books, err := s.store.GetAll()

	if err != nil {
		s.logger.Log("el error es &v", err.Error())
		return nil, err
	}

	return books, nil

}

func (s *Services) ObtenLibrosPorID(id int) (*model.Libro, error) {
	return s.store.GetById(id)
}

func (s *Services) CrearLibro(Libro model.Libro) (model.Libro, error) {

	if Libro.Title == "" {
		return model.Libro{}, errors.New("colocar titulo por favor")
	}

	return s.store.Create(Libro)
}

func (s *Services) RemoverLibro(id int) error {
	return s.store.Delete(id)
}
