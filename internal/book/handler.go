package book

import (
	"fmt"
	"gin-tutorial/config/res"
	"gin-tutorial/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler maneja las peticiones HTTP del recurso "books"
type Handler struct {
	repo Repository
}

// NewHandler crea un nuevo handler con un repositorio inyectado
func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// CreateBook maneja POST /books
func (h *Handler) CreateBook(context *gin.Context) {
	var book Book

	// Bind convierte el JSON del body a la estructura Book

	err := context.ShouldBindJSON(&book)

	if err != nil {
		res.Error(context, http.StatusBadRequest, err.Error())
		return
	}

	if book.Title == "" || book.Author == "" {
		res.Error(context, http.StatusBadRequest, "Llenar los campos")
		return
	}

	// Guardamos el libro
	if err := h.repo.Create(&book); err != nil {

		res.Error(context, http.StatusBadRequest, "No se pudo crear el libro")

	}

	context.JSON(http.StatusCreated, book)
}

// GetBooks maneja GET /books
func (h *Handler) GetBooks(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	fmt.Println(page)
	fmt.Println(size)
	booksPage, err := res.Paginate[Book](database.DB, page, size)
	if err != nil {
		res.Error(c, http.StatusInternalServerError, "No se pudo obtener libros")
		return
	}

	res.Success(c, "Libros obtenidos correctamente", booksPage)
}

// GetBookByID maneja GET /books/:id
func (h *Handler) GetBookByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	book, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Libro no encontrado"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook maneja PUT /books/:id
func (h *Handler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = uint(id)

	if err := h.repo.Update(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el libro"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook maneja DELETE /books/:id
func (h *Handler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el libro"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Libro eliminado"})
}
