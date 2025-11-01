package book

import (
	"fmt"
	"gin-tutorial/config/res"
	// "gin-tutorial/config/validator"
	"gin-tutorial/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handler maneja las peticiones HTTP del recurso "books"
type Handler struct {
	repo Repository
}

func NewService(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// CreateBook maneja POST /books
func (h *Handler) CreateBook(context *gin.Context) {

	var book Book

	err := context.ShouldBindJSON(&book)

	if err != nil {
		res.Error(context, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var validate = validator.New()

	if err := validate.Struct(book); err != nil {
		err := err.(validator.ValidationErrors)[0]
		res.Error(context, http.StatusBadRequest, fmt.Sprintf("El campo %v %v", err.Field(), err.Tag()), nil)
		return
	}

	if err := h.repo.Create(&book); err != nil {
		res.Error(context, http.StatusBadRequest, "No se pudo crear el libro", nil)
	}

	res.Created(context, "Creado con exito", book)
}

// GetBooks maneja GET /books
func (h *Handler) GetBooks(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	booksPage, err := res.Paginate[Book](database.DB, page, size)
	if err != nil {
		res.Error(c, http.StatusInternalServerError, "No se pudo obtener libros", nil)
		return
	}

	res.Success(c, "Libros obtenidos correctamente", booksPage)
}

// GetBookByID maneja GET /books/:id
func (h *Handler) GetBookByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Error(c, http.StatusInternalServerError, "No se pudo obtenido el id", nil)
		return
	}

	book, err := h.repo.FindByID(uint(id))
	if err != nil {
		res.Error(c, http.StatusInternalServerError, "No se pudo obtenido el book", nil)
		return
	}

	res.Success(c, fmt.Sprintf("Obtenido con el book con id %v", id), book)

}

// UpdateBook maneja PUT /books/:id
func (h *Handler) UpdateBook(c *gin.Context) {

	var book Book

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		res.Error(c, http.StatusInternalServerError, "No se pudo obtenido el id", nil)
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		res.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if book.Title == "" || book.Author == "" {
		res.Error(c, http.StatusInternalServerError, "No se pudo obtenido el title o author", nil)
		return
	}

	_, errById := h.repo.FindByID(uint(id))

	if errById != nil {
		res.Error(c, http.StatusInternalServerError, "No se pudo obtener el libre o no existe", nil)
		return
	}

	book.ID = uint(id)
	if err := h.repo.Update(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el libro"})
		return
	}
	res.Created(c, "Actualizado el libro", book)
}

// DeleteBook maneja DELETE /books/:id
func (h *Handler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Error(c, http.StatusInternalServerError, "No se pudo obtenido el id", nil)
		return
	}

	_, errById := h.repo.FindByID(uint(id))

	if err := h.repo.Delete(uint(id)); err != nil || errById != nil {
		res.Error(c, http.StatusInternalServerError, "No se pudo eliminar el libro o no existe", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Libro eliminado"})
}
